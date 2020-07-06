package calendar

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/soypita/otus_golang_homework/hw12_13_14_15_calendar/internal/models"
)

type Srv interface {
	CreateEvent(context.Context, *models.Event) (uuid.UUID, error)
	UpdateEvent(context.Context, uuid.UUID, *models.Event) error
	DeleteEvent(context.Context, uuid.UUID) error
	GetAllEvents(context.Context) ([]*models.Event, error)
	GetEventByID(context.Context, uuid.UUID) (*models.Event, error)
	FindDayEvents(context.Context, time.Time) ([]*models.Event, error)
	FindWeekEvents(context.Context, time.Time) ([]*models.Event, error)
	FindMonthEvents(context.Context, time.Time) ([]*models.Event, error)
}
