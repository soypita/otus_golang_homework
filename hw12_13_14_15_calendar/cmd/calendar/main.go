package main

import (
	"flag"
	"net"
	"os"

	"github.com/soypita/otus_golang_homework/hw12_13_14_15_calendar/internal/api/grpc"
	"github.com/soypita/otus_golang_homework/hw12_13_14_15_calendar/internal/api/rest"
	"github.com/soypita/otus_golang_homework/hw12_13_14_15_calendar/internal/configs"
	"github.com/soypita/otus_golang_homework/hw12_13_14_15_calendar/internal/logger"
	"github.com/soypita/otus_golang_homework/hw12_13_14_15_calendar/internal/providers"
	"github.com/soypita/otus_golang_homework/hw12_13_14_15_calendar/internal/services/calendar"
)

func main() {
	configPath := flag.String("config", "configs/config.yml", "path to config file")
	flag.Parse()
	config, err := configs.NewConfig(*configPath)
	if err != nil {
		panic(err)
	}

	logFile, err := os.OpenFile(config.Log.Path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}

	log, err := logger.NewLogger(logFile, config.Log.Level)
	if err != nil {
		panic(err)
	}

	restAddr := net.JoinHostPort(config.Host, config.RestPort)
	grpcAddr := net.JoinHostPort(config.Host, config.GrpcPort)

	repo, err := providers.NewRepository(log, config.Database.DSN, config.Database.InMemory)
	if err != nil {
		log.Fatalf("failed to initialize repository %s", err)
	}
	calendar := calendar.NewCalendar(repo)

	grpcServer := grpc.NewCalendarAPIServer(log, grpcAddr, calendar)
	restServer := rest.NewCalendarAPIServer(log, restAddr, calendar)

	if err := grpcServer.Start(); err != nil {
		log.Fatalf("failed to start grpc server %s", err)
	}

	if err := restServer.ListenAndServe(); err != nil {
		log.Fatalf("failed to stop rest server %s", err)
	}

	grpcServer.Stop()
}
