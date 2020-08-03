FROM golang:1.14-alpine AS build_base

RUN apk add --no-cache git

RUN mkdir -p /tmp/scheduler
WORKDIR /tmp/scheduler

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o /tmp/scheduler/scheduler cmd/calendarscheduler/main.go

FROM alpine:latest

COPY --from=build_base /tmp/scheduler/scheduler /app/scheduler

RUN mkdir /app/logs
WORKDIR /app

CMD ["./scheduler", "--config=config.yml"]