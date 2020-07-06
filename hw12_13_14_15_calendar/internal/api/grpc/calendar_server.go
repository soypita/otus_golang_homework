//go:generate protoc -I ../../../api/proto/ events_service.proto --go_out=plugins=grpc:.
package grpc

import (
	"context"
	"fmt"
	"net"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/uuid"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/sirupsen/logrus"
	"github.com/soypita/otus_golang_homework/hw12_13_14_15_calendar/internal/models"
	"github.com/soypita/otus_golang_homework/hw12_13_14_15_calendar/internal/services/calendar"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type CalendarAPIServer struct {
	log             logrus.FieldLogger
	host            string
	grcServer       *grpc.Server
	calendarService calendar.Srv
	interceptors    []grpc.UnaryServerInterceptor
}

func NewCalendarAPIServer(log logrus.FieldLogger, host string, calendarService calendar.Srv, interceptors []grpc.UnaryServerInterceptor) *CalendarAPIServer {
	return &CalendarAPIServer{
		log:             log,
		host:            host,
		calendarService: calendarService,
		interceptors:    interceptors,
	}
}

func (c *CalendarAPIServer) Start() error {
	lis, err := net.Listen("tcp", c.host)
	if err != nil {
		return fmt.Errorf("failed to listen %w", err)
	}

	opt := make([]grpc.ServerOption, 0, len(c.interceptors)+1)
	opt = append(opt, grpc.UnaryInterceptor(grpc_logrus.UnaryServerInterceptor(logrus.NewEntry(c.log.(*logrus.Logger)))))

	for _, interceptor := range c.interceptors {
		opt = append(opt, grpc.UnaryInterceptor(interceptor))
	}

	grpcServer := grpc.NewServer(opt...)

	reflection.Register(grpcServer)

	RegisterCalendarServer(grpcServer, c)
	c.grcServer = grpcServer

	c.log.Printf("Starting grpc server on %s", c.host)
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			c.log.Fatalf("failed to start server %w", err)
		}
	}()

	c.log.Printf("Successfully start grpc server")
	return nil
}

func (c *CalendarAPIServer) Stop() {
	c.log.Printf("Stopping grpc server...")
	c.grcServer.GracefulStop()
	c.log.Printf("Successfully stop grpc server")
}

func (c *CalendarAPIServer) CreateEvent(ctx context.Context, ev *CreateEventRequest) (*CreateEventResponse, error) {
	mEv, err := unmarshalEvent(ev.Event)
	if err != nil {
		return nil, err
	}
	id, err := c.calendarService.CreateEvent(ctx, mEv)
	if err != nil {
		return nil, err
	}
	return &CreateEventResponse{
		Id: id.String(),
	}, nil
}

func (c *CalendarAPIServer) UpdateEvent(ctx context.Context, ev *EventUpdateRequest) (*empty.Empty, error) {
	id, err := uuid.Parse(ev.Id)
	if err != nil {
		return nil, err
	}
	mEv, err := unmarshalEvent(ev.Event)
	if err != nil {
		return nil, err
	}
	err = c.calendarService.UpdateEvent(ctx, id, mEv)
	if err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}

func (c *CalendarAPIServer) DeleteEvent(ctx context.Context, ev *DeleteEventRequest) (*empty.Empty, error) {
	id, err := uuid.Parse(ev.Id)
	if err != nil {
		return nil, err
	}
	err = c.calendarService.DeleteEvent(ctx, id)
	if err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}

func (c *CalendarAPIServer) GetAllEvents(ctx context.Context, em *empty.Empty) (*GetAllEventsResponse, error) {
	mEvs, err := c.calendarService.GetAllEvents(ctx)
	if err != nil {
		return nil, err
	}
	resEvs := make([]*Event, 0, len(mEvs))
	for _, mEv := range mEvs {
		ev, err := marshalEvent(mEv)
		if err != nil {
			return nil, err
		}
		resEvs = append(resEvs, ev)
	}
	return &GetAllEventsResponse{
		Events: resEvs,
	}, nil
}

func (c *CalendarAPIServer) GetEventByID(ctx context.Context, ev *GetEventByIDRequest) (*GetEventByIDResponse, error) {
	id, err := uuid.Parse(ev.Id)
	if err != nil {
		return nil, err
	}
	mEv, err := c.calendarService.GetEventByID(ctx, id)
	if err != nil {
		return nil, err
	}
	res, err := marshalEvent(mEv)
	if err != nil {
		return nil, err
	}
	return &GetEventByIDResponse{Event: res}, nil
}

func (c *CalendarAPIServer) FindDayEvents(ctx context.Context, dayDate *FindDayEventsRequest) (*FindDayEventsResponse, error) {
	day, err := ptypes.Timestamp(dayDate.Date)
	if err != nil {
		return nil, fmt.Errorf("error while unmarshal day data %w", err)
	}
	mEvs, err := c.calendarService.FindDayEvents(ctx, day)
	if err != nil {
		return nil, err
	}
	resEvs := make([]*Event, 0, len(mEvs))
	for _, mEv := range mEvs {
		ev, err := marshalEvent(mEv)
		if err != nil {
			return nil, err
		}
		resEvs = append(resEvs, ev)
	}
	return &FindDayEventsResponse{
		Events: resEvs,
	}, nil
}

func (c *CalendarAPIServer) FindWeekEvents(ctx context.Context, weekDay *FindWeekEventsRequest) (*FindWeekEventsResponse, error) {
	day, err := ptypes.Timestamp(weekDay.Date)
	if err != nil {
		return nil, fmt.Errorf("error while unmarshal day data %w", err)
	}
	mEvs, err := c.calendarService.FindWeekEvents(ctx, day)
	if err != nil {
		return nil, err
	}
	resEvs := make([]*Event, 0, len(mEvs))
	for _, mEv := range mEvs {
		ev, err := marshalEvent(mEv)
		if err != nil {
			return nil, err
		}
		resEvs = append(resEvs, ev)
	}
	return &FindWeekEventsResponse{
		Events: resEvs,
	}, nil
}

func (c *CalendarAPIServer) FindMonthEvents(ctx context.Context, monthDay *FindMonthEventsRequest) (*FindMonthEventsResponse, error) {
	day, err := ptypes.Timestamp(monthDay.Date)
	if err != nil {
		return nil, fmt.Errorf("error while unmarshal day data %w", err)
	}
	mEvs, err := c.calendarService.FindMonthEvents(ctx, day)
	if err != nil {
		return nil, err
	}
	resEvs := make([]*Event, 0, len(mEvs))
	for _, mEv := range mEvs {
		ev, err := marshalEvent(mEv)
		if err != nil {
			return nil, err
		}
		resEvs = append(resEvs, ev)
	}
	return &FindMonthEventsResponse{
		Events: resEvs,
	}, nil
}

func unmarshalEvent(ev *Event) (*models.Event, error) {
	date, err := ptypes.Timestamp(ev.Date)
	if err != nil {
		return nil, fmt.Errorf("failed to umarshall event date : %w", err)
	}
	duration, err := ptypes.Duration(ev.Duration)
	if err != nil {
		return nil, fmt.Errorf("failed to umarshall event duration : %w", err)
	}
	notifyBefore, err := ptypes.Duration(ev.NotifyBefore)
	if err != nil {
		return nil, fmt.Errorf("failed to umarshall event notifyBefore : %w", err)
	}

	return &models.Event{
		Header:       ev.Header,
		Date:         date,
		Duration:     duration,
		Description:  ev.Description,
		OwnerID:      uuid.MustParse(ev.OwnerId),
		NotifyBefore: notifyBefore,
	}, nil
}

func marshalEvent(mEv *models.Event) (*Event, error) {
	date, err := ptypes.TimestampProto(mEv.Date)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal event date : %w", err)
	}
	duration := ptypes.DurationProto(mEv.Duration)

	notifyBefore := ptypes.DurationProto(mEv.NotifyBefore)

	return &Event{
		Header:       mEv.Header,
		Date:         date,
		Duration:     duration,
		Description:  mEv.Description,
		OwnerId:      mEv.OwnerID.String(),
		NotifyBefore: notifyBefore,
	}, nil
}
