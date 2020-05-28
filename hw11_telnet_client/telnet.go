package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

const (
	dialErrMsg      = "error to dial to host "
	connectedMsg    = "...Connected to"
	connectCloseMsg = "...Connection was closed by peer"
	eofMsg          = "...EOF"
)

type TelnetClient interface {
	Connect() error
	Close() error
	Receive() error
	Send() error
}

func NewTelnetClient(address string, timeout time.Duration, in io.ReadCloser, out io.Writer) TelnetClient {
	res := &TelnetClientImpl{
		address: address,
		timeout: timeout,
		in:      in,
		out:     out,
		doneCh:  make(chan struct{}),
		sendCh:  make(chan struct{}),
	}

	return res
}

type TelnetClientImpl struct {
	address string
	timeout time.Duration
	doneCh  chan struct{}
	sendCh  chan struct{}
	conn    net.Conn
	in      io.ReadCloser
	out     io.Writer
}

func (t *TelnetClientImpl) Connect() error {
	conn, err := net.DialTimeout("tcp", t.address, t.timeout)
	if err != nil {
		return fmt.Errorf("%s : %w", dialErrMsg, err)
	}
	t.conn = conn
	fmt.Fprintf(os.Stderr, "%s %s\n", connectedMsg, t.address)
	return nil
}

func (t *TelnetClientImpl) Receive() error {
	scanner := bufio.NewScanner(t.conn)
	for {
		select {
		case <-t.doneCh:
			return nil
		default:
			if !scanner.Scan() {
				return nil
			}
			fmt.Fprintln(t.out, scanner.Text())
		}
	}
}

func (t *TelnetClientImpl) Send() error {
	defer close(t.sendCh)
	scanner := bufio.NewScanner(t.in)
	for scanner.Scan() {
		str := scanner.Text() + "\n"
		_, err := t.conn.Write([]byte(str))
		if err != nil {
			fmt.Fprintln(os.Stderr, connectCloseMsg)
			return nil
		}
	}
	fmt.Fprintln(os.Stderr, eofMsg)
	return nil
}

func (t *TelnetClientImpl) Close() error {
	<-t.sendCh
	close(t.doneCh)
	err := t.conn.Close()
	if err != nil {
		return err
	}
	return nil
}
