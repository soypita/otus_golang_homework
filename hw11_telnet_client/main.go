package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	timeoutVar := flag.String("timeout", "10s", "connection timeout for client in seconds")
	flag.Parse()
	args := flag.Args()
	if err := runProcess(timeoutVar, args); err != nil {
		log.Fatalf(err.Error())
	}
}

func runProcess(timeoutVar *string, args []string) error {
	timeoutDurationStr := strings.TrimSuffix(*timeoutVar, "s")
	timeoutDuration, err := strconv.ParseInt(timeoutDurationStr, 10, 64)
	if err != nil {
		return fmt.Errorf("parse timeout error : %w", err)
	}

	if len(args) != 2 {
		return fmt.Errorf("usage: go-telnet --timeout=10s host port")
	}
	host := args[0]
	port := args[1]

	client := NewTelnetClient(net.JoinHostPort(host, port), time.Duration(timeoutDuration)*time.Second, os.Stdin, os.Stdout)

	if err := client.Connect(); err != nil {
		return err
	}

	errorCh := make(chan error, 1)
	go func() {
		errorCh <- client.Send()
	}()
	go func() {
		errorCh <- client.Receive()
	}()

	if err := client.Close(); err != nil {
		return err
	}
	if err = <-errorCh; err != nil {
		return err
	}
	return nil
}
