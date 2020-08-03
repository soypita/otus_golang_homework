package providers

import (
	"context"
	"fmt"
	"time"

	"github.com/cenkalti/backoff/v3"
	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose"
	"github.com/soypita/otus_golang_homework/hw12_13_14_15_calendar/internal/repository/pg"

	"github.com/soypita/otus_golang_homework/hw12_13_14_15_calendar/internal/repository"
	"github.com/soypita/otus_golang_homework/hw12_13_14_15_calendar/internal/repository/inmemory"

	_ "github.com/lib/pq" // Postgres driver
	"github.com/sirupsen/logrus"
)

func NewRepository(log logrus.FieldLogger, dsn string, isInMemory bool) (repository.EventsRepository, error) {
	if isInMemory {
		return inmemory.NewInMemRepository(log), nil
	}

	// try to connect with backoff
	be := backoff.NewExponentialBackOff()
	be.MaxElapsedTime = 10 * time.Second
	be.InitialInterval = 1 * time.Second
	be.MaxInterval = 5 * time.Second

	b := backoff.WithContext(be, context.Background())
	for {
		d := b.NextBackOff()
		if d == backoff.Stop {
			return nil, fmt.Errorf("error while connect to database. stop reconnecting")
		}

		<-time.After(d)
		db, err := sqlx.Connect("postgres", dsn)
		if err != nil {
			log.Println("error while connect to database. reconnect...")
			continue
		}

		if err := db.Ping(); err != nil {
			return nil, fmt.Errorf("error while connect to repository %w", err)
		}

		pgRep := pg.NewPGRepository(log, db)
		return pgRep, nil
	}
}

func RunMigration(repo repository.EventsRepository, dir string) error {
	// run migration
	db := repo.GetDB()
	if db == nil {
		return fmt.Errorf("cannot run migration")
	}
	if err := goose.Run("up", db.DB, dir); err != nil {
		return fmt.Errorf("could not run migration: %w", err)
	}
	return nil
}
