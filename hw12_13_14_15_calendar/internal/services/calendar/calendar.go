package calendar

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/soypita/otus_golang_homework/hw12_13_14_15_calendar/internal/models"
	"github.com/soypita/otus_golang_homework/hw12_13_14_15_calendar/internal/repository"
)

type Calendar struct {
	events repository.EventsRepository
}

func NewCalendar(repo repository.EventsRepository) *Calendar {
	return &Calendar{events: repo}
}

func (c *Calendar) CreateEvent(ev *models.Event) (uuid.UUID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	ev.ID = uuid.New()
	return c.events.CreateEvent(ctx, ev)
}

func (c *Calendar) UpdateEvent(id uuid.UUID, ev *models.Event) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return c.events.UpdateEvent(ctx, id, ev)
}

func (c *Calendar) DeleteEvent(id uuid.UUID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return c.events.DeleteEvent(ctx, id)
}

func (c *Calendar) GetAllEvents() ([]*models.Event, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return c.events.GetAllEvents(ctx)
}

func (c *Calendar) GetEventByID(id uuid.UUID) (*models.Event, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return c.events.GetEventByID(ctx, id)
}

func (c *Calendar) FindDayEvents(startDay time.Time) ([]*models.Event, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return c.events.FindDayEvents(ctx, startDay)
}

func (c *Calendar) FindWeekEvents(weekStartDay time.Time) ([]*models.Event, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return c.events.FindWeekEvents(ctx, weekStartDay)
}

func (c *Calendar) FindMonthEvents(monthStartDay time.Time) ([]*models.Event, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return c.events.FindMonthEvents(ctx, monthStartDay)
}
