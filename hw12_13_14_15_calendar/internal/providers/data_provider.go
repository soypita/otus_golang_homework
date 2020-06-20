package providers

import (
	"fmt"
	"github.com/soypita/otus_golang_homework/hw12_13_14_15_calendar/internal/repository/inmemory"
	"github.com/soypita/otus_golang_homework/hw12_13_14_15_calendar/internal/repository/pg"

	"github.com/soypita/otus_golang_homework/hw12_13_14_15_calendar/internal/repository"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // Postgres driver
	"github.com/sirupsen/logrus"
)

func NewRepository(log logrus.FieldLogger, dsn string, isInMemory bool) (repository.EventsRepository, error) {
	if isInMemory {
		return inmemory.NewInMemRepository(log), nil
	}
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("error while connect to database %w", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error while connect to repository %w", err)
	}
	pgRep := pg.NewPGRepository(log, db)
	return pgRep, nil
}
