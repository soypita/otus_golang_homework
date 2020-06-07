package repository

import (
	"context"
	"github/soypita/otus_golang_homework/hw12_13_14_15_calendar/internal/models"
	"time"

	"github.com/google/uuid"
)

type EventsRepository interface {
	CreateEvent(context.Context, *models.Event) (*models.Event, error)
	UpdateEvent(context.Context, uuid.UUID, *models.Event) error
	DeleteEvent(context.Context, uuid.UUID) error
	GetAllEvents(context.Context) ([]*models.Event, error)
	GetEventByID(context.Context, uuid.UUID) (*models.Event, error)
	FindDayEvents(context.Context, time.Time) ([]*models.Event, error)
	FindWeekEvents(context.Context, time.Time) ([]*models.Event, error)
	FindMonthEvents(context.Context, time.Time) ([]*models.Event, error)
}

type ErrDateBusy struct {
}

func (e ErrDateBusy) Error() string {
	return "current date already busy"
}

type ErrEventNotFound struct {
}

func (e ErrEventNotFound) Error() string {
	return "event not found in repository"
}
