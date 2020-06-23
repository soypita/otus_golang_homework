package inmemory

import (
	"context"
	"sync"
	"time"

	"github.com/soypita/otus_golang_homework/hw12_13_14_15_calendar/internal/repository"

	"github.com/soypita/otus_golang_homework/hw12_13_14_15_calendar/internal/models"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type InMemRepository struct {
	log    logrus.FieldLogger
	mtx    *sync.RWMutex
	events map[uuid.UUID]*models.Event
}

func NewInMemRepository(log logrus.FieldLogger) repository.EventsRepository {
	return &InMemRepository{
		log:    log,
		mtx:    &sync.RWMutex{},
		events: make(map[uuid.UUID]*models.Event),
	}
}

func (r *InMemRepository) CreateEvent(ctx context.Context, event *models.Event) (uuid.UUID, error) {
	r.mtx.RLock()
	for _, e := range r.events {
		if event.Date.Equal(e.Date) {
			return uuid.Nil, repository.ErrDateBusy{}
		}
	}
	r.mtx.RUnlock()

	r.mtx.Lock()
	r.events[event.ID] = event
	r.mtx.Unlock()
	return event.ID, nil
}

func (r *InMemRepository) UpdateEvent(ctx context.Context, id uuid.UUID, event *models.Event) error {
	r.mtx.RLock()
	for _, e := range r.events {
		if event.Date.Equal(e.Date) {
			return repository.ErrDateBusy{}
		}
	}
	r.mtx.RUnlock()

	r.mtx.Lock()
	r.events[id] = event
	r.mtx.Unlock()
	return nil
}

func (r *InMemRepository) DeleteEvent(ctx context.Context, id uuid.UUID) error {
	r.mtx.Lock()
	delete(r.events, id)
	r.mtx.Unlock()
	return nil
}

func (r *InMemRepository) GetAllEvents(ctx context.Context) ([]*models.Event, error) {
	res := make([]*models.Event, 0, len(r.events))
	r.mtx.RLock()
	for _, event := range r.events {
		res = append(res, event)
	}
	r.mtx.RUnlock()
	return res, nil
}

func (r *InMemRepository) GetEventByID(ctx context.Context, id uuid.UUID) (*models.Event, error) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	ev, ok := r.events[id]
	if !ok {
		return nil, repository.ErrEventNotFound{}
	}
	return ev, nil
}

func (r *InMemRepository) FindDayEvents(ctx context.Context, day time.Time) ([]*models.Event, error) {
	var dayEvents []*models.Event
	r.mtx.RLock()
	for _, val := range r.events {
		if (val.Date.Equal(day) || val.Date.After(day)) && val.Date.Before(day.AddDate(0, 0, 1)) {
			dayEvents = append(dayEvents, val)
		}
	}
	r.mtx.RUnlock()
	return dayEvents, nil
}

func (r *InMemRepository) FindWeekEvents(ctx context.Context, dayWeek time.Time) ([]*models.Event, error) {
	var dayEvents []*models.Event
	r.mtx.RLock()
	for _, val := range r.events {
		if (val.Date.Equal(dayWeek) || val.Date.After(dayWeek)) && val.Date.Before(dayWeek.AddDate(0, 0, 7)) {
			dayEvents = append(dayEvents, val)
		}
	}
	r.mtx.RUnlock()
	return dayEvents, nil
}

func (r *InMemRepository) FindMonthEvents(ctx context.Context, dayMonth time.Time) ([]*models.Event, error) {
	var dayEvents []*models.Event
	r.mtx.RLock()
	for _, val := range r.events {
		if (val.Date.Equal(dayMonth) || val.Date.After(dayMonth)) && val.Date.Before(dayMonth.AddDate(0, 1, 0)) {
			dayEvents = append(dayEvents, val)
		}
	}
	r.mtx.RUnlock()
	return dayEvents, nil
}
