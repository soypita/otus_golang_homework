package main

import (
	"flag"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	api "github.com/soypita/otus_golang_homework/hw12_13_14_15_calendar/internal/api/grpc"
	"google.golang.org/grpc"

	"github.com/soypita/otus_golang_homework/hw12_13_14_15_calendar/internal/configs/schedulercfg"
	"github.com/soypita/otus_golang_homework/hw12_13_14_15_calendar/internal/logger"
	"github.com/soypita/otus_golang_homework/hw12_13_14_15_calendar/internal/pubsub/publisher/ampq"
	"github.com/soypita/otus_golang_homework/hw12_13_14_15_calendar/internal/services/calendarscheduler"
)

func main() {
	configPath := flag.String("config", "configs/scheduler/config.yml", "path to config file")
	flag.Parse()

	config, err := schedulercfg.NewConfig(*configPath)
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

	pub := ampq.NewPublisher(log, config.AMPQ.URI, config.AMPQ.ExchangeName, config.AMPQ.ExchangeType, config.AMPQ.QueueName)
	err = pub.Connect()
	if err != nil {
		log.Fatal(err)
	}

	conn, err := grpc.Dial(net.JoinHostPort(config.EventAPI.Host, config.EventAPI.GrpcPort), grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := api.NewCalendarClient(conn)

	notifyCh := make(chan os.Signal, 1)
	signal.Notify(notifyCh, syscall.SIGINT, syscall.SIGTERM)

	notifTick := time.NewTicker(time.Duration(config.Schedule.Notify) * time.Second)
	cleanupTick := time.NewTicker(time.Duration(config.Schedule.Clean) * time.Second)
	defer notifTick.Stop()
	defer cleanupTick.Stop()

	scheduler := calendarscheduler.NewSchedulerService(log, pub, client, notifTick, cleanupTick)

	go func() {
		<-notifyCh
		scheduler.Stop()
	}()

	scheduler.Start()
}
