package calendarscheduler

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	server "github.com/soypita/otus_golang_homework/hw12_13_14_15_calendar/internal/api/grpc"
	"github.com/soypita/otus_golang_homework/hw12_13_14_15_calendar/internal/models"
	"github.com/soypita/otus_golang_homework/hw12_13_14_15_calendar/internal/repository/inmemory"
	"github.com/soypita/otus_golang_homework/hw12_13_14_15_calendar/internal/services/calendar/simple"
	"github.com/soypita/otus_golang_homework/hw12_13_14_15_calendar/pkg/api"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"log"
	"net"
	"testing"
	"time"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

type mockPublisher struct {
	Msgs []*api.NotificationDTO
}

func (m mockPublisher) Connect() error {
	return nil
}

func (m *mockPublisher) Send(ev *api.NotificationDTO) error {
	log.Printf("send message : %v", ev)
	m.Msgs = append(m.Msgs, ev)
	return nil
}

func TestBasicScheduler(t *testing.T) {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	log := logrus.New()
	repo := inmemory.NewInMemRepository(log)
	c := simple.NewCalendar(repo)
	server.RegisterCalendarServer(s, server.NewCalendarAPIServer(log, "", c, nil))
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()

	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()

	client := server.NewCalendarClient(conn)
	mP := &mockPublisher{}
	scheduler := NewSchedulerService(log, mP, client)

	t.Run(`should successfully send day notification`, func(t *testing.T) {
		ev, err := marshalEvent(&models.Event{
			ID:           uuid.New(),
			Header:       "Test",
			Date:         time.Now().Add(1 * time.Hour),
			Duration:     time.Duration(5) * time.Hour,
			Description:  "Test event",
			OwnerID:      uuid.New(),
			NotifyBefore: 1 * time.Hour,
		})
		assert.NoError(t, err)

		resp, err := client.CreateEvent(ctx, &server.CreateEventRequest{Event: ev})
		assert.NoError(t, err)
		assert.NotEqual(t, uuid.Nil, resp.Id)

		err = scheduler.ProcessDayEvents()
		assert.NoError(t, err)
		assert.NotEmpty(t, mP.Msgs)
		sendEv := mP.Msgs[0]
		assert.Equal(t, resp.Id, sendEv.EventID.String())
	})

	t.Run(`should successfully delete old notification`, func(t *testing.T) {
		ev, err := marshalEvent(&models.Event{
			ID:           uuid.New(),
			Header:       "Test",
			Date:         time.Now().AddDate(-1, 0, 5),
			Duration:     time.Duration(5) * time.Hour,
			Description:  "Test event",
			OwnerID:      uuid.New(),
			NotifyBefore: 1 * time.Hour,
		})
		assert.NoError(t, err)

		resp, err := client.CreateEvent(ctx, &server.CreateEventRequest{Event: ev})
		assert.NoError(t, err)
		assert.NotEqual(t, uuid.Nil, resp.Id)
		err = scheduler.DeleteOldData()
		assert.NoError(t, err)
		_, err = client.GetEventByID(ctx, &server.GetEventByIDRequest{Id: resp.Id})
		assert.Error(t, err)
		assert.Equal(t, err.Error(), "rpc error: code = Unknown desc = event not found in repository")
	})
}

func marshalEvent(mEv *models.Event) (*server.Event, error) {
	date, err := ptypes.TimestampProto(mEv.Date)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal event date : %w", err)
	}
	duration := ptypes.DurationProto(mEv.Duration)

	notifyBefore := ptypes.DurationProto(mEv.NotifyBefore)

	return &server.Event{
		Id:           mEv.ID.String(),
		Header:       mEv.Header,
		Date:         date,
		Duration:     duration,
		Description:  mEv.Description,
		OwnerId:      mEv.OwnerID.String(),
		NotifyBefore: notifyBefore,
	}, nil
}
