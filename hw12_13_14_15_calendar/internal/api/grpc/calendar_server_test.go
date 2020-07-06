package grpc

import (
	"context"
	"github.com/golang/protobuf/ptypes/duration"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/soypita/otus_golang_homework/hw12_13_14_15_calendar/internal/providers"
	"github.com/soypita/otus_golang_homework/hw12_13_14_15_calendar/internal/services/calendar/simple"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"net"
	"testing"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func bufDialer(context.Context, string) (net.Conn, error) {
	return lis.Dial()
}

func TestBasicGrpcServer(t *testing.T) {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	log := logrus.New()
	repo, err := providers.NewRepository(log, "", true)
	if err != nil {
		log.Fatalf("failed to initialize repository %s", err)
	}
	c := simple.NewCalendar(repo)
	RegisterCalendarServer(s, &CalendarAPIServer{
		log:             log,
		calendarService: c,
	})
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
	client := NewCalendarClient(conn)

	t.Run("successfully create event", func(t *testing.T) {
		resp, err := client.CreateEvent(ctx, &CreateEventRequest{
			Event: &Event{Header: "Test",
				Date: &timestamp.Timestamp{
					Seconds: 100,
					Nanos:   0,
				},
				Duration: &duration.Duration{
					Seconds: 10,
					Nanos:   0,
				},
				Description: "Test",
				OwnerId:     uuid.New().String(),
				NotifyBefore: &duration.Duration{
					Seconds: 0,
					Nanos:   0,
				},
			},
		})
		assert.Nil(t, err)
		assert.NotNil(t, resp)
		assert.True(t, resp.Id != "")
	})

	t.Run("should return create event error", func(t *testing.T) {
		_, err := client.CreateEvent(ctx, &CreateEventRequest{
			Event: &Event{
				Header: "Test",
				Date: &timestamp.Timestamp{
					Seconds: 100,
					Nanos:   0,
				},
				Duration: &duration.Duration{
					Seconds: 10,
					Nanos:   0,
				},
				Description: "Test",
				OwnerId:     uuid.New().String(),
				NotifyBefore: &duration.Duration{
					Seconds: 0,
					Nanos:   0,
				},
			},
		})
		assert.NotNil(t, err)
	})
}
