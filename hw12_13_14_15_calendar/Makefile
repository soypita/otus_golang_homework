build:
	go build -o calendar cmd/calendar/calendar.go

test:
	go test -race ./...

lint:
	golangci-lint run ./...

run:
	docker build . -t calendar
	docker-compose up -d

migrations:
	goose -dir migrations/ postgres "host=localhost port=5432 user=soypita password=soypita dbname=calendar sslmode=disable" up

.PHONY: build
