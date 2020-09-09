
#build stage
FROM golang:alpine3.12

RUN apk add --no-cache bash git openssh

RUN apk add musl-dev gcc

WORKDIR /go/src/app


