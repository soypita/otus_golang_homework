package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/soypita/otus_golang_homework/hw12_13_14_15_calendar/internal/configs/sendercfg"
	"github.com/soypita/otus_golang_homework/hw12_13_14_15_calendar/internal/logger"
	"github.com/soypita/otus_golang_homework/hw12_13_14_15_calendar/internal/pubsub/subscriber/ampq"
	"github.com/soypita/otus_golang_homework/hw12_13_14_15_calendar/internal/services/calendarsender"
)

func main() {
	configPath := flag.String("config", "configs/sender/config.yml", "path to config file")
	flag.Parse()
	config, err := sendercfg.NewConfig(*configPath)
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

	sub := ampq.NewSubscriber(log, config.AMPQ.URI, config.AMPQ.ExchangeName, config.AMPQ.ExchangeType, config.AMPQ.QueueName)
	sender := calendarsender.NewSenderService(log, sub, &config.Sink)

	notifyCh := make(chan os.Signal, 1)
	signal.Notify(notifyCh, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		<-notifyCh
		cancel()
	}()

	log.Println("start to listen events queue...")
	if err := sender.ListenAndProcess(ctx); err != nil {
		log.Fatal(err)
	}

	log.Println("successfully stop listener")
}
