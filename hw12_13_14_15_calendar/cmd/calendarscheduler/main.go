package main

import (
	"flag"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"

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

	pub := ampq.NewPublisher(log, config.AMPQ.URI, config.AMPQ.ExchangeName, config.AMPQ.ExchangeType, config.AMPQ.QueueName)
	err = pub.Connect()
	if err != nil {
		panic(err)
	}

	conn, err := grpc.Dial(net.JoinHostPort(config.EventAPI.Host, config.EventAPI.GrpcPort), grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := api.NewCalendarClient(conn)

	scheduler := calendarscheduler.NewSchedulerService(log, pub, client)

	notifyCh := make(chan os.Signal, 1)
	signal.Notify(notifyCh, syscall.SIGINT, syscall.SIGTERM)

	notifTick := time.NewTicker(time.Duration(config.Schedule.Notify) * time.Second)
	cleanupTick := time.NewTicker(time.Duration(config.Schedule.Clean) * time.Second)

	go scheduleNotifications(notifTick, log, scheduler)
	go scheduleCleanup(cleanupTick, log, scheduler)

	<-notifyCh
	notifTick.Stop()
	cleanupTick.Stop()
	log.Println("scheduler successfully stop")
}

func scheduleNotifications(tick *time.Ticker, log *logrus.Logger, scheduler *calendarscheduler.SchedulerService) {
	for range tick.C {
		log.Println("start to process notifications...")
		if err := scheduler.ProcessDayEvents(); err != nil {
			log.Printf("error while running scheduler %s\n", err)
		}
	}
}

func scheduleCleanup(tick *time.Ticker, log *logrus.Logger, scheduler *calendarscheduler.SchedulerService) {
	for range tick.C {
		log.Println("start to cleanup...")
		if err := scheduler.DeleteOldData(); err != nil {
			log.Printf("error while running scheduler %s\n", err)
		}
	}
}
