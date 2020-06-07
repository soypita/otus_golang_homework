package main

import (
	"context"
	"flag"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/soypita/otus_golang_homework/hw12_13_14_15_calendar/internal/configs"
	"github.com/soypita/otus_golang_homework/hw12_13_14_15_calendar/internal/handlers"
	"github.com/soypita/otus_golang_homework/hw12_13_14_15_calendar/internal/logger"

	"github.com/gorilla/mux"
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

	h := handlers.NewEventsHandler(log)
	router := mux.NewRouter()
	router.Use(h.LoggingMiddleware)

	router.HandleFunc("/health", h.HealthCheck).Methods("GET")

	addr := net.JoinHostPort(config.Host, config.Port)
	server := http.Server{
		Addr:         addr,
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	notifyCh := make(chan os.Signal, 1)
	errorCh := make(chan error)
	signal.Notify(notifyCh, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-notifyCh
		log.Println("Stopping server")
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		server.SetKeepAlivesEnabled(false)
		if err := server.Shutdown(ctx); err != nil {
			errorCh <- err
		}
		close(errorCh)
	}()

	log.Printf("Start server on %s....", addr)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Panic(err)
	}
	if err := <-errorCh; err != nil {
		log.Panic(err)
	}
	log.Println("Stop server successfully")
}
