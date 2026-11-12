#!/usr/bin/env bash
set -euo pipefail
cd "$(dirname "$0")/.."
export AWS_DEFAULT_REGION="${AWS_REGION:-us-east-2}" AWS_PAGER=""
host=${HOSTNAME_FOR_SITE:-www.go.chrisfoderingham.com}
stack=go-path-student
repo=go-path-student
: "${REVIEWER_GITHUB_IDS:?Set numeric IDs of the independent course reviewers}"
: "${CERTIFICATE_ARN:?Set the ARN of an issued ACM certificate covering the site hostname}"
secret=$(aws secretsmanager describe-secret --secret-id "${OAUTH_SECRET_ID:-go-path/github-oauth}" --query ARN --output text)
status=$(aws acm describe-certificate --certificate-arn "$CERTIFICATE_ARN" --query Certificate.Status --output text)
[[ $status == ISSUED ]] || { echo "ACM certificate is not issued."; exit 1; }
vpc=${VPC_ID:-$(aws ec2 describe-vpcs --filters Name=is-default,Values=true --query 'Vpcs[0].VpcId' --output text)}
if [[ -z ${SUBNET_A:-} || -z ${SUBNET_B:-} ]]; then
  mapfile -t defaults < <(aws ec2 describe-subnets --filters Name=vpc-id,Values="$vpc" Name=default-for-az,Values=true --query 'Subnets[].[AvailabilityZone,SubnetId]' --output text | sort | awk '!seen[$1]++ {print $2}')
  [[ ${#defaults[@]} -ge 2 ]] || { echo "Supply VPC_ID, SUBNET_A and SUBNET_B for two public subnets in separate zones."; exit 1; }
  SUBNET_A=${defaults[0]}; SUBNET_B=${defaults[1]}
fi
aws cloudformation validate-template --template-body file://deploy/aws.yml >/dev/null
if [[ -n $(git status --porcelain --untracked-files=no) ]]; then echo "Deploy from a clean checkout."; exit 1; fi
sha=$(git rev-parse HEAD)
account=$(aws sts get-caller-identity --query Account --output text)
registry="$account.dkr.ecr.$AWS_DEFAULT_REGION.amazonaws.com"
if ! aws ecr describe-repositories --repository-names "$repo" >/dev/null 2>&1; then
  aws ecr create-repository --repository-name "$repo" --image-tag-mutability IMMUTABLE --image-scanning-configuration scanOnPush=true >/dev/null
fi
if ! aws ecr describe-images --repository-name "$repo" --image-ids imageTag="$sha" >/dev/null 2>&1; then
  aws ecr get-login-password | docker login --username AWS --password-stdin "$registry"
  docker build --platform linux/amd64 --tag "$registry/$repo:$sha" .
  docker push "$registry/$repo:$sha"
fi
digest=$(aws ecr describe-images --repository-name "$repo" --image-ids imageTag="$sha" --query 'imageDetails[0].imageDigest' --output text)
aws cloudformation deploy --stack-name "$stack" --template-file deploy/aws.yml --capabilities CAPABILITY_IAM --no-fail-on-empty-changeset --parameter-overrides Vpc="$vpc" SubnetA="$SUBNET_A" SubnetB="$SUBNET_B" Image="$registry/$repo@$digest" CertificateArn="$CERTIFICATE_ARN" OAuthSecretArn="$secret" Hostname="$host" Reviewers="$REVIEWER_GITHUB_IDS"
alb=$(aws cloudformation describe-stacks --stack-name "$stack" --query "Stacks[0].Outputs[?OutputKey=='LoadBalancerDNS'].OutputValue" --output text)
route=(--connect-to "$host:443:$alb:443" --retry 5 --retry-all-errors)
curl "${route[@]}" -fsS "https://$host/healthz"
code=$(curl "${route[@]}" -sS -o /dev/null -w '%{http_code}' "https://$host/student/state")
[[ $code == 401 ]] || { echo "Expected unauthenticated student state to return 401; got $code."; exit 1; }
# A healthy catalog can still lack its student OAuth configuration.
login=$(curl "${route[@]}" -sS -o /dev/null -w '%{http_code} %{redirect_url}' "https://$host/student/login")
[[ $login == "302 https://github.com/login/oauth/authorize?"* ]] || { echo "Student sign-in did not start GitHub authorization."; exit 1; }
echo "Verified $sha on $alb. Route 53 alias target for $host: $alb"

