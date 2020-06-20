package pg

import (
	"context"
	"fmt"
	"github.com/soypita/otus_golang_homework/hw12_13_14_15_calendar/internal/repository"
	"strings"
	"time"

	"github.com/soypita/otus_golang_homework/hw12_13_14_15_calendar/internal/models"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	pq "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

type PGRepository struct {
	log logrus.FieldLogger
	db  *sqlx.DB
}

func NewPGRepository(log logrus.FieldLogger, db *sqlx.DB) repository.EventsRepository {
	return &PGRepository{
		log: log,
		db:  db,
	}
}

func (r *PGRepository) CreateEvent(ctx context.Context, event *models.Event) (uuid.UUID, error) {
	_, err := r.db.NamedExecContext(ctx,
		"INSERT INTO events (id, header, date, duration, description, ownerid, notifybefore) VALUES (:id, :header, :date, :duration, :description, :ownerid, :notifybefore)",
		event)
	if err != nil {
		if resCode, ok := err.(*pq.Error); ok {
			if resCode.Code == "23505" {
				return uuid.Nil, repository.ErrDateBusy{}
			}
		}
		return uuid.Nil, fmt.Errorf("error while create new event: %w", err)
	}
	return event.ID, nil
}

func (r *PGRepository) UpdateEvent(ctx context.Context, id uuid.UUID, event *models.Event) error {
	result, err := r.db.NamedExecContext(ctx,
		"UPDATE events SET header = :header, date = :date, duration = :duration, description = :description, ownerid = :ownerid, notifybefore = :notifybefore WHERE id = :id AND date != :date",
		map[string]interface{}{
			"id":           id,
			"header":       event.Header,
			"date":         event.Date,
			"duration":     event.Duration,
			"description":  event.Description,
			"ownerid":      event.OwnerID,
			"notifybefore": event.NotifyBefore,
		})
	if err != nil {
		return fmt.Errorf("error while update record with id %d : %w", id, err)
	}
	count, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error while update record with id %d : %w", id, err)
	}
	if count == 0 {
		return repository.ErrDateBusy{}
	}
	return nil
}

func (r *PGRepository) DeleteEvent(ctx context.Context, id uuid.UUID) error {
	_, err := r.db.NamedExecContext(ctx,
		"DELETE FROM events WHERE id = :id",
		map[string]interface{}{
			"id": id,
		})
	if err != nil {
		return fmt.Errorf("error while delete record with id %d : %w", id, err)
	}
	return err
}

func (r *PGRepository) GetAllEvents(ctx context.Context) ([]*models.Event, error) {
	var events []*models.Event
	err := r.db.SelectContext(ctx, &events,
		`SELECT * FROM events`)
	if err != nil {
		return nil, fmt.Errorf("error while get all events : %w", err)
	}
	return events, nil
}

func (r *PGRepository) GetEventByID(ctx context.Context, id uuid.UUID) (*models.Event, error) {
	event := models.Event{}
	err := r.db.GetContext(ctx, &event,
		`SELECT * FROM events WHERE id = $1`, id)
	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return nil, repository.ErrEventNotFound{}
		}
		return nil, fmt.Errorf("error while find day events : %w", err)
	}
	return &event, nil
}

func (r *PGRepository) FindDayEvents(ctx context.Context, day time.Time) ([]*models.Event, error) {
	var events []*models.Event
	err := r.db.SelectContext(ctx, &events,
		`SELECT * FROM events WHERE date BETWEEN $1 AND $1 + (interval '1d')`, day)
	if err != nil {
		return nil, fmt.Errorf("error while find day events : %w", err)
	}
	return events, nil
}

func (r *PGRepository) FindWeekEvents(ctx context.Context, day time.Time) ([]*models.Event, error) {
	var events []*models.Event
	err := r.db.SelectContext(ctx, &events,
		`SELECT * FROM events WHERE date BETWEEN $1 AND $1 + (interval '7 weeks')`, day)
	if err != nil {
		return nil, fmt.Errorf("error while find day events : %w", err)
	}
	return events, nil
}

func (r PGRepository) FindMonthEvents(ctx context.Context, day time.Time) ([]*models.Event, error) {
	var events []*models.Event
	err := r.db.SelectContext(ctx, &events,
		`SELECT * FROM events WHERE date BETWEEN $1 AND $1 + (interval '1 months')`, day)
	if err != nil {
		return nil, fmt.Errorf("error while find day events : %w", err)
	}
	return events, nil
}
