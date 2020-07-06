package rest

import (
	"github.com/sirupsen/logrus"
	"github.com/soypita/otus_golang_homework/hw12_13_14_15_calendar/internal/providers"
	"github.com/soypita/otus_golang_homework/hw12_13_14_15_calendar/internal/services/calendar/simple"
	"github.com/steinfletcher/apitest"
	"net/http"
	"testing"
)

func TestBasicRestServer(t *testing.T) {
	log := logrus.New()
	repo, err := providers.NewRepository(log, "", true)
	if err != nil {
		log.Fatalf("failed to initialize repository %s", err)
	}
	c := simple.NewCalendar(repo)
	server := CalendarAPIServer{
		calendarService: c,
		log:             log,
	}

	t.Run("successfully create event", func(t *testing.T) {
		apitest.New().
			HandlerFunc(server.CreateEvent).
			Put("/events").
			JSONFromFile("../../../tests/testdata/rest/create_event.json").
			Expect(t).
			Status(http.StatusOK).
			End()
	})
	t.Run("should return create event error", func(t *testing.T) {
		apitest.New().
			HandlerFunc(server.CreateEvent).
			Put("/events").
			JSONFromFile("../../../tests/testdata/rest/create_event.json").
			Expect(t).
			Status(http.StatusInternalServerError).
			End()
	})
}
