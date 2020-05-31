package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	timeoutVar := flag.Duration("timeout", 10*time.Second, "connection timeout for client in seconds")
	flag.Parse()
	args := flag.Args()
	if err := runProcess(timeoutVar, args); err != nil {
		log.Fatalf(err.Error())
	}
}

func runProcess(timeoutVar *time.Duration, args []string) (err error) {
	if len(args) != 2 {
		err = fmt.Errorf("usage: go-telnet --timeout=10s host port")
		return
	}
	host := args[0]
	port := args[1]

	client := NewTelnetClient(net.JoinHostPort(host, port), *timeoutVar, os.Stdin, os.Stdout)
	if err = client.Connect(); err != nil {
		return err
	}
	defer func() {
		err = client.Close()
	}()

	notifyCh := make(chan os.Signal, 1)
	errorCh := make(chan error, 1)
	signal.Notify(notifyCh, syscall.SIGINT, syscall.SIGTERM)

	go processReceive(&client, errorCh)
	go processSend(&client, errorCh)

	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		for {
			select {
			case <-notifyCh:
				return
			case err = <-errorCh:
				if err != nil {
					if err == io.EOF {
						_, err = fmt.Fprintln(os.Stderr, eofMsg)
					}
					return
				}
			default:
				continue
			}
		}
	}()

	wg.Wait()
	return
}

func processSend(client *TelnetClient, errorCh chan error) {
	for {
		err := (*client).Send()
		if err != nil {
			errorCh <- err
			return
		}
	}
}

func processReceive(client *TelnetClient, errorCh chan error) {
	for {
		err := (*client).Receive()
		if err != nil {
			errorCh <- err
			return
		}
	}
}
