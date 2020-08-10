package main

import (
	"flag"
	"log"
	"net"
	"os"

	"github.com/soypita/otus_golang_homework/hw12_13_14_15_calendar/internal/configs/calendarcfg"

	"github.com/soypita/otus_golang_homework/hw12_13_14_15_calendar/internal/services/calendar/simple"

	"github.com/soypita/otus_golang_homework/hw12_13_14_15_calendar/internal/api/grpc"
	"github.com/soypita/otus_golang_homework/hw12_13_14_15_calendar/internal/api/rest"
	"github.com/soypita/otus_golang_homework/hw12_13_14_15_calendar/internal/logger"
	"github.com/soypita/otus_golang_homework/hw12_13_14_15_calendar/internal/providers"
)

func main() {
	configPath := flag.String("config", "configs/calendar/config.yml", "path to config file")
	flag.Parse()
	config, err := calendarcfg.NewConfig(*configPath)
	if err != nil {
		log.Fatal(err)
	}

	logFile, err := os.OpenFile(config.Log.Path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}

	log, err := logger.NewLogger(logFile, config.Log.Level)
	if err != nil {
		log.Fatal(err)
	}

	restAddr := net.JoinHostPort(config.Host, config.RestPort)
	grpcAddr := net.JoinHostPort(config.Host, config.GrpcPort)

	repo, err := providers.NewRepository(log, config.Database.DSN, config.Database.InMemory)
	if err != nil {
		log.Fatalf("failed to initialize repository %s", err)
	}
	calendar := simple.NewCalendar(repo)

	grpcServer := grpc.NewCalendarAPIServer(log, grpcAddr, calendar, nil)
	restServer := rest.NewCalendarAPIServer(log, restAddr, calendar, nil)

	if err := grpcServer.Start(); err != nil {
		log.Fatalf("failed to start grpc server %s", err)
	}

	if err := restServer.ListenAndServe(); err != nil {
		log.Fatalf("failed to stop rest server %s", err)
	}

	grpcServer.Stop()
}
