build:
	go build -o bin/calendar cmd/calendar/main.go

test:
	go test -race ./...

install-deps:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint

lint: install-deps
	golangci-lint run ./...

run:
	docker build . -t calendar
	docker-compose up -d

migrations:
	goose -dir migrations/ postgres "host=localhost port=5432 user=soypita password=soypita dbname=calendar sslmode=disable" up

.PHONY: build

generate:
	go generate ./...