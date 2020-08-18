FROM golang:1.14-alpine AS build_base

RUN apk add --no-cache git

RUN mkdir -p /tmp/sender
WORKDIR /tmp/sender

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o /tmp/sender/sender cmd/calendarsender/main.go

FROM alpine:latest

COPY --from=build_base /tmp/sender/sender /app/sender

RUN mkdir /app/logs
WORKDIR /app

CMD ["./sender", "--config=config.yml"]