FROM golang:1.10.2-alpine3.7 AS build
RUN apk --no-cache add gcc g++ make ca-certificates
WORKDIR /go/src/github.com/ravjotsingh9/discussionForum-Web-Service

COPY Gopkg.lock Gopkg.toml ./
COPY vendor vendor
COPY web-service web-service
COPY util util
COPY db db
COPY schema schema

RUN go install ./...

FROM alpine:3.7
WORKDIR /usr/bin
COPY --from=build /go/bin .
