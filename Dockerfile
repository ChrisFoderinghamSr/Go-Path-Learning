FROM golang:1.26-alpine AS build
WORKDIR /src
COPY go.mod *.go units.json ./
COPY dist/ dist/
RUN CGO_ENABLED=0 go test ./... && CGO_ENABLED=0 go build -trimpath -o /go-path .
RUN mkdir /data && chown 65532:65532 /data && apk add --no-cache ca-certificates
FROM scratch
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build --chown=65532:65532 /data /data
ENV STUDENT_DATA_DIR=/data
COPY --from=build /go-path /go-path
USER 65532:65532
ENV GO_PATH_ADDR=:8080
EXPOSE 8080
ENTRYPOINT ["/go-path"]

