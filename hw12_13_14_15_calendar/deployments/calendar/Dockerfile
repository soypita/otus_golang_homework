FROM golang:1.14-alpine AS build_base

RUN apk add --no-cache git

RUN mkdir -p /tmp/calendar

WORKDIR /tmp/calendar

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o /tmp/calendar/calendar cmd/calendar/main.go

FROM alpine:latest

COPY --from=build_base /tmp/calendar/calendar /app/calendar

RUN mkdir /app/logs
WORKDIR /app

CMD ["./calendar", "--config=config.yml"]