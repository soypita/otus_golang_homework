package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"github.com/soypita/otus_golang_homework/hw12_13_14_15_calendar/pkg/api"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/cucumber/godog"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/duration"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/google/uuid"
	grpc_api "github.com/soypita/otus_golang_homework/hw12_13_14_15_calendar/internal/api/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CalendarFeature struct {
	client             grpc_api.CalendarClient
	createEventResp    *grpc_api.CreateEventResponse
	findEventsSize     int
	responseStatusCode int
	responseBody       []byte
	grpcStsCode        codes.Code
}

func (c *CalendarFeature) initCalendarClient(*godog.Scenario) {
	endpoint := os.Getenv("CALENDAR_ENDPOINT")
	if endpoint == "" {
		endpoint = ":8090"
	}
	conn, err := grpc.Dial(endpoint, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect to calendar service: %s", err.Error())
	}
	c.client = grpc_api.NewCalendarClient(conn)
}

func (c *CalendarFeature) clenup(sc *godog.Scenario, err error) {
	resp, err := c.client.GetAllEvents(context.Background(), &empty.Empty{})
	if err != nil {
		log.Fatalf("failed to clenup: %s", err.Error())
	}
	for _, e := range resp.Events {
		c.client.DeleteEvent(context.Background(), &grpc_api.DeleteEventRequest{
			Id: e.Id,
		})
	}
	c.createEventResp = nil
	c.responseStatusCode = 0
	c.responseBody = nil
	c.grpcStsCode = 0
	c.findEventsSize = 0
}

func (c *CalendarFeature) iSendRequestTo(httpMethod, addr string) (err error) {
	var r *http.Response

	switch httpMethod {
	case http.MethodGet:
		r, err = http.Get(addr)
	default:
		err = fmt.Errorf("unknown method: %s", httpMethod)
	}

	if err != nil {
		return
	}
	c.responseStatusCode = r.StatusCode
	c.responseBody, err = ioutil.ReadAll(r.Body)
	r.Body.Close()
	return
}

func (c *CalendarFeature) responseCodeShouldBe(code int) error {
	if c.responseStatusCode != code {
		return fmt.Errorf("unexpected status code: %d != %d", c.responseStatusCode, code)
	}
	return nil
}

func (c *CalendarFeature) statusInResponseShouldBe(srvSts string) error {
	resp := make(map[string]string)
	if err := json.Unmarshal(c.responseBody, &resp); err != nil {
		return fmt.Errorf("failed to parse response: %w", err)
	}
	sts, ok := resp["status"]
	if !ok {
		return fmt.Errorf("no status field in response")
	}
	if sts != srvSts {
		return fmt.Errorf("unexpected status in response: %s != %s", sts, srvSts)
	}
	return nil
}

func (c *CalendarFeature) iCallCreateEventMethod() error {
	resp, err := c.client.CreateEvent(context.Background(), &grpc_api.CreateEventRequest{
		Event: &grpc_api.Event{
			Id:     uuid.New().String(),
			Header: "Test",
			Date: &timestamp.Timestamp{
				Seconds: 1000,
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
	if err != nil {
		return fmt.Errorf("failed to create event: %w", err)
	}
	c.createEventResp = resp
	c.grpcStsCode = status.Code(err)
	return nil
}

func (c *CalendarFeature) responseShouldHaveEventID() error {
	if c.createEventResp == nil {
		return fmt.Errorf("create event response is nil")
	}

	if c.createEventResp.Id == "" {
		return fmt.Errorf("create event response ID is blank")
	}

	if c.createEventResp.Id == uuid.Nil.String() {
		return fmt.Errorf("create event response ID is nil uuid")
	}

	return nil
}

func (c *CalendarFeature) iCallCreateEventMethodForDate(date string) error {
	tms, err := c.parseDateToTimestamp(date)
	if err != nil {
		return err
	}
	resp, err := c.client.CreateEvent(context.Background(), &grpc_api.CreateEventRequest{
		Event: &grpc_api.Event{
			Id:     uuid.New().String(),
			Header: "Test",
			Date:   tms,
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
	c.createEventResp = resp
	c.grpcStsCode = status.Code(err)
	return nil
}

func (c *CalendarFeature) iCallCreateEventMethodForCurrentDate() error {
	resp, err := c.client.CreateEvent(context.Background(), &grpc_api.CreateEventRequest{
		Event: &grpc_api.Event{
			Id:     uuid.New().String(),
			Header: "Test notifications",
			Date:   ptypes.TimestampNow(),
			Duration: &duration.Duration{
				Seconds: 10,
				Nanos:   0,
			},
			Description: "Test notifications",
			OwnerId:     uuid.New().String(),
			NotifyBefore: &duration.Duration{
				Seconds: 0,
				Nanos:   0,
			},
		},
	})
	c.createEventResp = resp
	c.grpcStsCode = status.Code(err)
	return nil
}

func (c *CalendarFeature) iCallCreateEventMethodForOldDate() error {
	oldDate := time.Now().AddDate(-1, 0, 1)
	tms, err := ptypes.TimestampProto(oldDate)
	if err != nil {
		return fmt.Errorf("failed to parse timestamp: %w", err)
	}
	resp, err := c.client.CreateEvent(context.Background(), &grpc_api.CreateEventRequest{
		Event: &grpc_api.Event{
			Id:     uuid.New().String(),
			Header: "Test cleanup event",
			Date:   tms,
			Duration: &duration.Duration{
				Seconds: 10,
				Nanos:   0,
			},
			Description: "Test cleanup event",
			OwnerId:     uuid.New().String(),
			NotifyBefore: &duration.Duration{
				Seconds: 0,
				Nanos:   0,
			},
		},
	})
	c.createEventResp = resp
	c.grpcStsCode = status.Code(err)
	return nil
}

func (c *CalendarFeature) eventsShouldClearAsync() error {
	time.Sleep(10 * time.Second)
	_, err := c.client.GetEventByID(context.Background(), &grpc_api.GetEventByIDRequest{Id: c.createEventResp.Id})
	if err == nil {
		return fmt.Errorf("should failed to get non existing event")
	}
	return nil
}

func (c *CalendarFeature) iShouldReceiveEventNotification() error {
	// Need to wait async notification event
	time.Sleep(5 * time.Second)

	logFile, err := os.OpenFile("out.txt", os.O_RDWR, 0644)
	defer logFile.Close()
	if err != nil {
		return fmt.Errorf("failed to read out file: %w", err)
	}
	reader := bufio.NewReader(logFile)

	notificationB, err := reader.ReadBytes('\n')
	if err != nil {
		return fmt.Errorf("failed to read line: %w", err)
	}
	var resultNotification api.NotificationDTO
	err = json.Unmarshal(notificationB, &resultNotification)
	if err != nil {
		return fmt.Errorf("failed to parser notification: %w", err)
	}
	err = logFile.Truncate(0)
	if err != nil {
		return fmt.Errorf("failed to tracate file: %w", err)
	}
	_, err = logFile.Seek(0, 0)
	if err != nil {
		return fmt.Errorf("failed to seek file: %w", err)
	}
	return nil
}

func (c *CalendarFeature) iGetErrorResponse() error {
	if c.grpcStsCode == codes.OK {
		return fmt.Errorf("error code should't be OK")
	}
	return nil
}

func (c *CalendarFeature) thereIsEventWithDate(date string) error {
	tms, err := c.parseDateToTimestamp(date)
	if err != nil {
		return err
	}
	resp, err := c.client.CreateEvent(context.Background(), &grpc_api.CreateEventRequest{
		Event: &grpc_api.Event{
			Id:     uuid.New().String(),
			Header: "Test",
			Date:   tms,
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
	if err != nil {
		return fmt.Errorf("failed to prepeare event: %w", err)
	}

	c.createEventResp = resp
	return nil
}

func (c *CalendarFeature) iCallUpdateEventMethod() error {
	_, err := c.client.UpdateEvent(context.Background(), &grpc_api.EventUpdateRequest{
		Id: c.createEventResp.Id,
		Event: &grpc_api.Event{
			Id:     uuid.New().String(),
			Header: "Update Test",
			Date:   ptypes.TimestampNow(),
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
	c.grpcStsCode = status.Code(err)
	return nil
}

func (c *CalendarFeature) iCallUpdateEventMethodFor(id string) error {
	_, err := c.client.UpdateEvent(context.Background(), &grpc_api.EventUpdateRequest{
		Id: id,
		Event: &grpc_api.Event{
			Id:     uuid.New().String(),
			Header: "Update Test",
			Date:   ptypes.TimestampNow(),
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
	c.grpcStsCode = status.Code(err)
	return nil
}

func (c *CalendarFeature) iGetSuccessResponse() error {
	if c.grpcStsCode != codes.OK {
		return fmt.Errorf("error code should be OK")
	}
	return nil
}

func (c *CalendarFeature) iCallDeleteEventMethod() error {
	_, err := c.client.DeleteEvent(context.Background(), &grpc_api.DeleteEventRequest{
		Id: c.createEventResp.Id,
	})
	c.grpcStsCode = status.Code(err)
	return nil
}

func (c *CalendarFeature) iCallDeleteEventMethodFor(id string) error {
	_, err := c.client.DeleteEvent(context.Background(), &grpc_api.DeleteEventRequest{
		Id: id,
	})
	c.grpcStsCode = status.Code(err)
	return nil
}

func (c *CalendarFeature) iCallFindDayEventsMethodForDay(day string) error {
	tms, err := c.parseDateToTimestamp(day)
	if err != nil {
		return err
	}
	resp, err := c.client.FindDayEvents(context.Background(), &grpc_api.FindDayEventsRequest{
		Date: tms,
	})

	c.grpcStsCode = status.Code(err)
	c.findEventsSize = len(resp.Events)
	return nil
}

func (c *CalendarFeature) eventsResponseSizeShouldBe(size int) error {
	if c.findEventsSize != size {
		return fmt.Errorf("wrong size of response: %d != %d", c.findEventsSize, size)
	}
	return nil
}

func (c *CalendarFeature) iCallFindWeekEventsMethodForDay(day string) error {
	tms, err := c.parseDateToTimestamp(day)
	if err != nil {
		return err
	}
	resp, err := c.client.FindWeekEvents(context.Background(), &grpc_api.FindWeekEventsRequest{
		Date: tms,
	})

	c.grpcStsCode = status.Code(err)
	c.findEventsSize = len(resp.Events)
	return nil
}

func (c *CalendarFeature) iCallFindMonthEventsMethodForDay(day string) error {
	tms, err := c.parseDateToTimestamp(day)
	if err != nil {
		return err
	}
	resp, err := c.client.FindMonthEvents(context.Background(), &grpc_api.FindMonthEventsRequest{
		Date: tms,
	})

	c.grpcStsCode = status.Code(err)
	c.findEventsSize = len(resp.Events)
	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	calendar := &CalendarFeature{}

	ctx.BeforeScenario(calendar.initCalendarClient)
	ctx.AfterScenario(calendar.clenup)

	ctx.Step(`^I send "([^"]*)" request to "([^"]*)"$`, calendar.iSendRequestTo)
	ctx.Step(`^Response code should be (\d+)$`, calendar.responseCodeShouldBe)
	ctx.Step(`^Status in response should be "([^"]*)"$`, calendar.statusInResponseShouldBe)
	ctx.Step(`^I call createEvent method$`, calendar.iCallCreateEventMethod)
	ctx.Step(`^Response should have event ID$`, calendar.responseShouldHaveEventID)
	ctx.Step(`^I call createEvent method for date "([^"]*)"$`, calendar.iCallCreateEventMethodForDate)
	ctx.Step(`^I get error response$`, calendar.iGetErrorResponse)
	ctx.Step(`^there is event with date "([^"]*)"$`, calendar.thereIsEventWithDate)
	ctx.Step(`^I call updateEvent method$`, calendar.iCallUpdateEventMethod)
	ctx.Step(`^I get success response$`, calendar.iGetSuccessResponse)
	ctx.Step(`^I call deleteEvent method$`, calendar.iCallDeleteEventMethod)
	ctx.Step(`^I call updateEvent method for "([^"]*)"$`, calendar.iCallUpdateEventMethodFor)
	ctx.Step(`^I call deleteEvent method for "([^"]*)"$`, calendar.iCallDeleteEventMethodFor)
	ctx.Step(`^Events response size should be (\d+)$`, calendar.eventsResponseSizeShouldBe)
	ctx.Step(`^I call findDayEvents method for day "([^"]*)"$`, calendar.iCallFindDayEventsMethodForDay)
	ctx.Step(`^I call findWeekEvents method for day "([^"]*)"$`, calendar.iCallFindWeekEventsMethodForDay)
	ctx.Step(`^I call findMonthEvents method for day "([^"]*)"$`, calendar.iCallFindMonthEventsMethodForDay)
	ctx.Step(`^I call createEvent method for current date$`, calendar.iCallCreateEventMethodForCurrentDate)
	ctx.Step(`^I should receive event notification$`, calendar.iShouldReceiveEventNotification)
	ctx.Step(`^Events should clear async$`, calendar.eventsShouldClearAsync)
	ctx.Step(`^I call createEvent method for old date$`, calendar.iCallCreateEventMethodForOldDate)

}

func (c *CalendarFeature) parseDateToTimestamp(date string) (*timestamp.Timestamp, error) {
	tm, err := time.Parse(time.RFC3339, date)
	if err != nil {
		return nil, fmt.Errorf("failed to parse date: %w", err)
	}
	tms, err := ptypes.TimestampProto(tm)
	if err != nil {
		return nil, fmt.Errorf("failed to parse date to timestamp: %w", err)
	}
	return tms, nil
}
