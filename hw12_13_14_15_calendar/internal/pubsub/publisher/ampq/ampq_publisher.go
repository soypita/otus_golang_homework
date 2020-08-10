package ampq

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/cenkalti/backoff/v3"
	"github.com/sirupsen/logrus"
	"github.com/soypita/otus_golang_homework/hw12_13_14_15_calendar/pkg/api"
	"github.com/streadway/amqp"
)

type Publisher struct {
	log          logrus.FieldLogger
	conn         *amqp.Connection
	channel      *amqp.Channel
	uri          string
	exchangeName string
	exchangeType string
	queue        string
	done         chan error
}

func NewPublisher(log logrus.FieldLogger, uri, exchangeName, exchangeType, queue string) *Publisher {
	return &Publisher{
		log:          log,
		uri:          uri,
		exchangeName: exchangeName,
		exchangeType: exchangeType,
		queue:        queue,
		done:         make(chan error),
	}
}

func (s *Publisher) reconnect(ctx context.Context) error {
	be := backoff.NewExponentialBackOff()
	be.MaxElapsedTime = time.Minute
	be.InitialInterval = 1 * time.Second
	be.Multiplier = 2
	be.MaxInterval = 15 * time.Second

	b := backoff.WithContext(be, ctx)
	for {
		d := b.NextBackOff()
		if d == backoff.Stop {
			return fmt.Errorf("stop reconnecting")
		}
		<-time.After(d)
		if err := s.connect(); err != nil {
			s.log.Printf("could not connect in reconnect call: %+v", err)
			continue
		}
		return nil
	}
}

func (s *Publisher) connect() error {
	var err error
	s.log.Println("try to connect to broker")
	s.conn, err = amqp.Dial(s.uri)
	if err != nil {
		return fmt.Errorf("error while dial: %s", err)
	}
	s.log.Println("success connect to broker")

	s.log.Println("try to get channel")
	s.channel, err = s.conn.Channel()
	if err != nil {
		return fmt.Errorf("error to get broker channel: %s", err)
	}
	s.log.Println("success get channel")

	go func() {
		s.log.Printf("closing: %s", <-s.conn.NotifyClose(make(chan *amqp.Error)))
		// Понимаем, что канал сообщений закрыт, надо пересоздать соединение.
		s.done <- errors.New("channel closed")
	}()

	s.log.Println("try to exchange declare")
	if err := s.channel.ExchangeDeclare(
		s.exchangeName, // name
		s.exchangeType, // type
		true,           // durable
		false,          // auto-deleted
		false,          // internal
		false,          // noWait
		nil,            // arguments
	); err != nil {
		return fmt.Errorf("error exchange declare: %s", err)
	}
	s.log.Println("success exchange declare")

	_, err = s.channel.QueueDeclare(
		s.queue, // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	if err != nil {
		return fmt.Errorf("error queue declare: %s", err)
	}

	return nil
}

func (s *Publisher) Connect(ctx context.Context) error {
	var err error
	if err = s.connect(); err != nil {
		return fmt.Errorf("error: %v", err)
	}

	go func() {
		for {
			select {
			case <-ctx.Done():
				s.log.Println(ctx.Err())
				return
			case resDone := <-s.done:
				if resDone != nil {
					s.log.Println("try to reconnect...")
					err := s.reconnect(ctx)
					if err != nil {
						s.log.Println("reconnecting error: ", err)
						return
					}
					s.log.Println("success reconnect")
				}
			default:
				continue
			}
		}
	}()
	return nil
}

func (s *Publisher) Send(ev *api.NotificationDTO) error {
	s.log.Println("try to send message...")
	body, err := json.Marshal(ev)
	if err != nil {
		return fmt.Errorf("error while marshal event entity %w", err)
	}
	err = s.channel.Publish(
		s.exchangeName,
		s.queue,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		return fmt.Errorf("error while send data to broker %w", err)
	}
	s.log.Println("success send message...")

	return nil
}
