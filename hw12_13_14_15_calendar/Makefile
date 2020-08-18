build:
	go build -o bin/calendar cmd/calendar/main.go
	go build -o bin/calendar_scheduler cmd/calendarscheduler/main.go
	go build -o bin/calendar_sender cmd/calendarsender/main.go

test:
	go test -race ./...

install-deps:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint

lint: install-deps
	golangci-lint run ./...

migrations:
	goose -dir migrations/ postgres "host=localhost port=5432 user=soypita password=soypita dbname=calendar sslmode=disable" up

.PHONY: build

generate:
	go generate ./...

up:
	docker-compose up -d --build

down:
	docker-compose down

bdd:
	set -e ;\
	docker-compose -f docker-compose.test.yml up -d ;\
	test_status_code=0 ;\
	chmod +x test_integration.sh ;\
	./test_integration.sh || test_status_code=$$? ;\
    docker-compose -f docker-compose.test.yml down ;\
    exit $$test_status_code ;\
