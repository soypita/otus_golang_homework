package calendarsender

import (
	"runtime"

	"github.com/sirupsen/logrus"
	"github.com/soypita/otus_golang_homework/hw12_13_14_15_calendar/internal/pubsub/subscriber"
	"github.com/soypita/otus_golang_homework/hw12_13_14_15_calendar/pkg/api"
)

type SenderService struct {
	log logrus.FieldLogger
	sub subscriber.Srv
}

func NewSenderService(log logrus.FieldLogger, pub subscriber.Srv) *SenderService {
	return &SenderService{
		log: log,
		sub: pub,
	}
}

func (s *SenderService) ListenAndProcess() error {
	h := func(msg *api.NotificationDTO) error {
		s.log.Printf("receive msg: %v\n", *msg)
		return nil
	}
	tNum := runtime.GOMAXPROCS(0)

	return s.sub.Listen(h, tNum)
}
