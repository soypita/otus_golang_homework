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
	docker-compose build && docker-compose up -d

down:
	docker-compose down

# Calendar
# docker build -t calendar:0.1.0 -f deployments/calendar/Dockerfile .
# docker run -v $(pwd)/configs/calendar/config.yml:/app/config.yml -p 8080:8080 calendar:0.1.0

# Scheduler
# docker build -t scheduler:0.1.0 -f deployments/scheduler/Dockerfile .
# docker run -v $(pwd)/configs/scheduler/config.yml:/app/config.yml scheduler:0.1.0

# Sender
# docker build -t sender:0.1.0 -f deployments/sender/Dockerfile .
# docker run -v $(pwd)/configs/sender/config.yml:/app/config.yml sender:0.1.0