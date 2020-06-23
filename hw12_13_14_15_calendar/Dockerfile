FROM golang:1.13-alpine AS build_base

RUN apk add --no-cache git

WORKDIR /tmp/calendar

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o calendar cmd/calendar/calendar.go

FROM alpine:3.9
RUN apk add ca-certificates

COPY --from=build_base /tmp/calendar/calendar /app/calendar