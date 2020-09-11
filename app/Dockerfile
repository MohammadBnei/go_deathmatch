
#build stage
FROM golang:alpine3.12

RUN apk add --no-cache bash git openssh

RUN apk add musl-dev gcc

RUN apk add apache2-utils

WORKDIR /go/src/app

COPY app/go.mod .
COPY app/go.sum .

RUN go mod download

RUN go get github.com/pilu/fresh

CMD go install && fresh


