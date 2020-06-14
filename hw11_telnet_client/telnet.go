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

func NewTelnetClient(address string, timeout time.Duration, in io.Reader, out io.Writer) TelnetClient {
	inputReader := bufio.NewReader(in)
	res := &TelnetClientImpl{
		address:     address,
		timeout:     timeout,
		inputReader: inputReader,
		out:         out,
	}

	return res
}

type TelnetClientImpl struct {
	address     string
	timeout     time.Duration
	conn        net.Conn
	inputReader *bufio.Reader
	connReader  *bufio.Reader
	out         io.Writer
}

func (t *TelnetClientImpl) Connect() error {
	conn, err := net.DialTimeout("tcp", t.address, t.timeout)
	if err != nil {
		return fmt.Errorf("%s : %w", dialErrMsg, err)
	}
	t.conn = conn
	t.connReader = bufio.NewReader(t.conn)
	fmt.Fprintf(os.Stderr, "%s %s\n", connectedMsg, t.address)
	return nil
}

func (t *TelnetClientImpl) Receive() error {
	line, err := t.connReader.ReadString('\n')
	if err != nil {
		if err == io.EOF {
			err = fmt.Errorf(connectCloseMsg)
		}
		return err
	}
	if _, err := fmt.Fprint(t.out, line); err != nil {
		return err
	}
	return nil
}

func (t *TelnetClientImpl) Send() error {
	line, err := t.inputReader.ReadString('\n')
	if err != nil {
		return err
	}
	_, err = t.conn.Write([]byte(line))
	if err != nil {
		return fmt.Errorf(connectCloseMsg)
	}
	return nil
}

func (t *TelnetClientImpl) Close() error {
	err := t.conn.Close()
	if err != nil {
		return err
	}
	return nil
}
