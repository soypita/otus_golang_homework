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

type Subscriber struct {
	log          logrus.FieldLogger
	conn         *amqp.Connection
	channel      *amqp.Channel
	uri          string
	exchangeName string
	exchangeType string
	queue        string
	done         chan error
}

func NewSubscriber(log logrus.FieldLogger, uri, exchangeName, exchangeType, queue string) *Subscriber {
	return &Subscriber{
		log:          log,
		uri:          uri,
		exchangeName: exchangeName,
		exchangeType: exchangeType,
		queue:        queue,
		done:         make(chan error),
	}
}

func (s *Subscriber) reconnect(ctx context.Context) (<-chan amqp.Delivery, error) {
	be := backoff.NewExponentialBackOff()
	be.MaxElapsedTime = time.Minute
	be.InitialInterval = 1 * time.Second
	be.Multiplier = 2
	be.MaxInterval = 15 * time.Second
	var err error
	b := backoff.WithContext(be, ctx)
	for {
		d := b.NextBackOff()
		if d == backoff.Stop {
			return nil, fmt.Errorf("stop reconnecting")
		}

		<-time.After(d)
		if err = s.connect(); err != nil {
			s.log.Printf("could not connect in reconnect call: %+v", err)
			continue
		}
		msgs, err := s.announceQueue()
		if err != nil {
			fmt.Printf("Couldn't connect: %+v", err)
			continue
		}
		return msgs, nil
	}
}

func (s *Subscriber) announceQueue() (<-chan amqp.Delivery, error) {
	queue, err := s.channel.QueueDeclare(
		s.queue,
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return nil, fmt.Errorf("queue Declare: %s", err)
	}

	err = s.channel.Qos(50, 0, false)
	if err != nil {
		return nil, fmt.Errorf("error setting qos: %s", err)
	}

	if err = s.channel.QueueBind(
		queue.Name,
		queue.Name,
		s.exchangeName,
		false,
		nil,
	); err != nil {
		return nil, fmt.Errorf("queue Bind: %s", err)
	}

	msgs, err := s.channel.Consume(
		queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("queue Consume: %s", err)
	}

	return msgs, nil
}

func (s *Subscriber) connect() error {
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
	return nil
}

//nolint:gocognit
func (s *Subscriber) Listen(ctx context.Context, handler func(msg *api.NotificationDTO) error) error {
	var err error
	if err = s.connect(); err != nil {
		return fmt.Errorf("error: %v", err)
	}

	s.log.Println("try to announce queue")
	msgs, err := s.announceQueue()
	if err != nil {
		return fmt.Errorf("error to announce queue: %w", err)
	}
	s.log.Println("success  announce queue")

	for {
		go func() {
			for {
				select {
				case <-ctx.Done():
					s.log.Println("finish sending: ", ctx.Err())
					return
				case msg, ok := <-msgs:
					if !ok {
						return
					}
					notification := &api.NotificationDTO{}
					if err := json.Unmarshal(msg.Body, notification); err != nil {
						s.log.Println("error while read notification from queue : %s", err)
						continue
					}
					if err := handler(notification); err != nil {
						s.log.Println("error while handle message: %s", err)
					}
				}
			}
		}()

		select {
		case <-ctx.Done():
			s.log.Println(ctx.Err())
			return nil
		case resDone := <-s.done:
			if resDone != nil {
				s.log.Println("try to reconnect...")
				msgs, err = s.reconnect(ctx)
				if err != nil {
					return fmt.Errorf("reconnecting error: %w", err)
				}
				s.log.Println("success reconnect")
			}
		}
	}
}
