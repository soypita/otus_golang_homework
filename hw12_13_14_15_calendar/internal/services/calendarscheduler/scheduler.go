package calendarscheduler

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/soypita/otus_golang_homework/hw12_13_14_15_calendar/internal/api/grpc"
	"github.com/soypita/otus_golang_homework/hw12_13_14_15_calendar/internal/pubsub/publisher"
	"github.com/soypita/otus_golang_homework/hw12_13_14_15_calendar/pkg/api"
)

type SchedulerService struct {
	log        logrus.FieldLogger
	pub        publisher.Srv
	client     grpc.CalendarClient
	notifyTick *time.Ticker
	cleanTick  *time.Ticker
	doneCh     chan struct{}
}

func NewSchedulerService(log logrus.FieldLogger, pub publisher.Srv, client grpc.CalendarClient, notifyTick, cleanTick *time.Ticker) *SchedulerService {
	return &SchedulerService{
		log:        log,
		pub:        pub,
		client:     client,
		notifyTick: notifyTick,
		cleanTick:  cleanTick,
		doneCh:     make(chan struct{}),
	}
}

func (s *SchedulerService) Start() {
	for {
		select {
		case <-s.doneCh:
			s.log.Println("scheduler successfully stop")
			return
		case <-s.notifyTick.C:
			s.log.Println("start to process notifications...")
			if err := s.processDayEvents(); err != nil {
				s.log.Printf("error while running scheduler %s\n", err)
			}
		case <-s.cleanTick.C:
			s.log.Println("start to cleanup...")
			if err := s.deleteOldData(); err != nil {
				s.log.Printf("error while running scheduler %s\n", err)
			}
		default:
			continue
		}
	}
}

func (s *SchedulerService) Stop() {
	close(s.doneCh)
}

func (s *SchedulerService) processDayEvents() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	ev, err := s.client.FindDayEvents(ctx, &grpc.FindDayEventsRequest{
		Date: ptypes.TimestampNow(),
	})

	if err != nil {
		return fmt.Errorf("error while pull day events %w", err)
	}

	if ev == nil {
		s.log.Println("no events to process")
		return nil
	}

	currTimestamp := time.Now()
	w := &sync.WaitGroup{}
	for _, event := range ev.Events {
		w.Add(1)
		go func(w *sync.WaitGroup, e *grpc.Event) {
			defer w.Done()
			mEv, err := unmarshalEvent(e)
			if err != nil {
				s.log.Printf("error while unmarshal event : %s\n", err)
			}

			notifyBefore, err := ptypes.Duration(e.NotifyBefore)
			if err != nil {
				s.log.Printf("error while parse notify before field: %s\n", err)
				return
			}
			notifyTimestamp := mEv.EventDate.Add(-notifyBefore)
			if currTimestamp.After(notifyTimestamp) || currTimestamp.Equal(notifyTimestamp) {
				err := s.pub.Send(mEv)
				if err != nil {
					s.log.Printf("error while sent event : %s\n", err)
				}
			}
		}(w, event)
	}
	w.Wait()
	return nil
}

func (s *SchedulerService) deleteOldData() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	currTms := time.Now().AddDate(-1, 0, 0)
	d, err := ptypes.TimestampProto(currTms)
	if err != nil {
		return fmt.Errorf("error while ger year timestamp %w", err)
	}
	ev, err := s.client.FindMonthEvents(ctx, &grpc.FindMonthEventsRequest{
		Date: d,
	})
	if err != nil {
		return fmt.Errorf("error while pull year events to delete %w", err)
	}

	if ev == nil {
		s.log.Println("no events to delete")
		return nil
	}

	w := &sync.WaitGroup{}
	for _, event := range ev.Events {
		w.Add(1)
		go func(w *sync.WaitGroup, e *grpc.Event) {
			defer w.Done()
			_, err := s.client.DeleteEvent(ctx, &grpc.DeleteEventRequest{
				Id: e.Id,
			})
			if err != nil {
				s.log.Printf("failed to delete event with id %s : %s\n", e.Id, err)
			}
		}(w, event)
	}
	w.Wait()
	return nil
}

func unmarshalEvent(mEv *grpc.Event) (*api.NotificationDTO, error) {
	date, err := ptypes.Timestamp(mEv.Date)
	if err != nil {
		return nil, fmt.Errorf("error while convert date field : %w", err)
	}

	id, err := uuid.Parse(mEv.Id)

	if err != nil {
		return nil, fmt.Errorf("error while convert id field : %w", err)
	}

	ownerID, err := uuid.Parse(mEv.OwnerId)
	if err != nil {
		return nil, fmt.Errorf("error while convert ownerID field : %w", err)
	}

	return &api.NotificationDTO{
		EventID:      id,
		EventHeader:  mEv.Header,
		EventDate:    date,
		EventOwnerID: ownerID,
	}, nil
}
