package calendarsender

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/soypita/otus_golang_homework/hw12_13_14_15_calendar/internal/configs/sendercfg"

	"github.com/sirupsen/logrus"
	"github.com/soypita/otus_golang_homework/hw12_13_14_15_calendar/internal/pubsub/subscriber"
	"github.com/soypita/otus_golang_homework/hw12_13_14_15_calendar/pkg/api"
)

type SenderService struct {
	log     logrus.FieldLogger
	sub     subscriber.Srv
	out     *sendercfg.Output
	outFile *os.File
}

func NewSenderService(log logrus.FieldLogger, pub subscriber.Srv, out *sendercfg.Output) *SenderService {
	return &SenderService{
		log: log,
		sub: pub,
		out: out,
	}
}

func (s *SenderService) ListenAndProcess(ctx context.Context) error {
	defer func() {
		if s.outFile != nil {
			s.log.Println("close out sink file")
			if err := s.outFile.Close(); err != nil {
				s.log.Println("failed to close out sink file: %s", err.Error())
			}
		}
	}()
	var h func(msg *api.NotificationDTO) error
	switch s.out.Type {
	case "console":
		h = s.consoleSink
	case "file":
		f, err := s.prepareFileSink(s.out.Name)
		if err != nil {
			return err
		}
		h = f
	}
	return s.sub.Listen(ctx, h)
}

func (s *SenderService) prepareFileSink(fileName string) (func(msg *api.NotificationDTO) error, error) {
	logFile, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	s.outFile = logFile
	if err != nil {
		return nil, fmt.Errorf("failed to prepeare file sink: %w", err)
	}
	return func(msg *api.NotificationDTO) error {
		err := json.NewEncoder(s.outFile).Encode(*msg)
		if err != nil {
			return err
		}
		return nil
	}, nil
}

func (s *SenderService) consoleSink(msg *api.NotificationDTO) error {
	s.log.Printf("receive msg: %v\n", *msg)
	return nil
}
