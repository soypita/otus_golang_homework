package simple

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

func (c *Calendar) CreateEvent(ctx context.Context, ev *models.Event) (uuid.UUID, error) {
	ev.ID = uuid.New()
	return c.events.CreateEvent(ctx, ev)
}

func (c *Calendar) UpdateEvent(ctx context.Context, id uuid.UUID, ev *models.Event) error {
	return c.events.UpdateEvent(ctx, id, ev)
}

func (c *Calendar) DeleteEvent(ctx context.Context, id uuid.UUID) error {
	return c.events.DeleteEvent(ctx, id)
}

func (c *Calendar) GetAllEvents(ctx context.Context) ([]*models.Event, error) {
	return c.events.GetAllEvents(ctx)
}

func (c *Calendar) GetEventByID(ctx context.Context, id uuid.UUID) (*models.Event, error) {
	return c.events.GetEventByID(ctx, id)
}

func (c *Calendar) FindDayEvents(ctx context.Context, startDay time.Time) ([]*models.Event, error) {
	return c.events.FindDayEvents(ctx, startDay)
}

func (c *Calendar) FindWeekEvents(ctx context.Context, weekStartDay time.Time) ([]*models.Event, error) {
	return c.events.FindWeekEvents(ctx, weekStartDay)
}

func (c *Calendar) FindMonthEvents(ctx context.Context, monthStartDay time.Time) ([]*models.Event, error) {
	return c.events.FindMonthEvents(ctx, monthStartDay)
}
