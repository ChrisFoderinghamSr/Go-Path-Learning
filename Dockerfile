FROM golang:1.26-alpine AS build
WORKDIR /src
COPY go.mod main.go main_test.go ./
COPY dist/ dist/
RUN CGO_ENABLED=0 go test ./... && CGO_ENABLED=0 go build -trimpath -o /go-path .
FROM scratch
COPY --from=build /go-path /go-path
USER 65532:65532
ENV GO_PATH_ADDR=:8080
EXPOSE 8080
ENTRYPOINT ["/go-path"]
