# Deploy the GitHub student edition

The student edition serves the website and authenticated API from one HTTPS origin. The existing Amplify static build alone cannot serve the OAuth callback, private records, reviews, or certificates. Test the branch on a staging hostname before directing students to it.

1. Register a GitHub OAuth App belonging to the course operator. Set its callback URL to https://YOUR-STAGING-HOST/student/callback and homepage to the same HTTPS origin. The app requests public_repo for writing milestones. GitHub issues this broader public repository permission; the server confines writes to the enrolled Go-Path-Student repository. Never put the client secret in frontend variables.
2. Build the Dockerfile and run the Go service on port 8080 behind HTTPS. An AWS deployment can use ECS with a load balancer and one running task. Provide an encrypted persistent filesystem at /data, writable by UID 65532. Configure backups and test restoring a student record. This edition requires a single writer. Stop the previous task before starting a replacement or migrate records to a transactional shared database before running multiple replicas.
3. Inject PUBLIC_ORIGIN (the external HTTPS origin, no trailing path), GITHUB_CLIENT_ID, GITHUB_CLIENT_SECRET (from Secrets Manager), STUDENT_DATA_DIR=/data, and REVIEWER_GITHUB_IDS (comma-separated numeric GitHub IDs of appointed instructors). Use /healthz for health checks. Session tokens are kept only in process memory; restarting requires sign-in again.
4. Test with a real student account and a different instructor account. Verify login, enrollment, opening and saving a lesson, restoring progress after sign-out, and a milestone commit to go-path/progress.json. Submit original coursework under work/UNIT-ID/, review the pinned diff and tests, and exercise both revision and approval paths. Confirm a student cannot review their own submissions. For delivery behind Amplify, proxy the whole protected application to this HTTPS service, or point the custom hostname to the service after staging passes.
5. Course completion requires all 40 lesson checkpoints and independent approval of 40 lesson submissions plus 12 projects. The instructor queue lives at /student. Reviewers check code, test results, assignment-specific allowed AI use and originality. A fork or known template repository is rejected. Local cloning and AI authorship cannot be reliably inferred from Git data. Ask for an explanation or demonstration when needed. The verification route on each issued certificate is shareable.

Visitor records contain dates only, counted once per browser cookie interval. A student's numeric GitHub ID owns their private record. Notes, drafts, quiz answers and reviewer feedback stay on the server. Public GitHub snapshots include lesson completion, last lesson, submission SHAs and status. Students can edit their repository; the private instructor record is authoritative. Students consent before enrollment. Publish a retention/deletion policy appropriate to your program and set filesystem backup retention before launch.

An existing browser-only learning backup can be imported by its owner after login. Imported checklist progress never grants instructor approval. To change course requirements, update units.json along with curriculum.js and review how existing students will meet new requirements.


## Launch from AWS CloudShell in Ohio

Use the repository's deploy/launch.sh from a clean checkout of the reviewed release commit. The script builds an immutable image in ECR, deploys the CloudFormation stack in us-east-2 and tests HTTPS using the load balancer before DNS cutover.

Account prerequisites:

- Create the course OAuth app in GitHub with homepage https://www.go.chrisfoderingham.com and callback https://www.go.chrisfoderingham.com/student/callback. Save GITHUB_CLIENT_ID and GITHUB_CLIENT_SECRET as two JSON keys in an Ohio Secrets Manager secret named go-path/github-oauth. Use the default Secrets Manager encryption key with this template. Keep the values out of GitHub and chat.
- Request and DNS-validate an ACM certificate covering www.go.chrisfoderingham.com in us-east-2. Set CERTIFICATE_ARN to the issued certificate ARN.
- Set REVIEWER_GITHUB_IDS to the appointed instructors' numeric GitHub IDs. Students cannot approve their own work. The repository owner's ID is 307307124 if the owner is an appointed instructor.
- Use a default VPC with two default public subnets or explicitly set VPC_ID, SUBNET_A and SUBNET_B to internet-routed subnets in separate availability zones with VPC DNS enabled.
- The deployment creates billable Fargate, ALB, EFS, ECR and CloudWatch resources. The ECS task role for infrastructure execution reads only the selected OAuth secret in addition to the standard ECS execution policy. Updating the single-writer task briefly interrupts service.

Run in an authenticated CloudShell from the repository root:

```bash
export CERTIFICATE_ARN='YOUR_ISSUED_ACM_CERTIFICATE_ARN'
export REVIEWER_GITHUB_IDS='YOUR_APPOINTED_REVIEWER_ID'
bash deploy/launch.sh
```

The script prints the load balancer DNS name after checking /healthz and confirming anonymous requests to /student/state receive 401. Create a Route 53 A alias for www.go pointing to that load balancer, then complete the real student and independent instructor test above. Existing go.chrisfoderingham.com can continue serving the Amplify edition during the transition. To roll back a service revision, check out a previous tested release and rerun the script with the same stack and volume. Retain the encrypted EFS filesystem and verify a backup restore separately.
