FROM golang:latest

RUN apt-get update
RUN apt-get upgrade -y

ENV GOBIN /go/bin

COPY migrations /migrations

RUN go get -u github.com/pressly/goose/cmd/goose