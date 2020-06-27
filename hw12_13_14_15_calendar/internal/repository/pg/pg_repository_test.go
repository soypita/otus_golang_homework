package pg

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose"
	"github.com/sirupsen/logrus"
	"github.com/soypita/otus_golang_homework/hw12_13_14_15_calendar/internal/repository"
	"github.com/soypita/otus_golang_homework/hw12_13_14_15_calendar/pkg/models"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"os"
	"sync"
	"testing"
	"time"

	_ "github.com/lib/pq" // Postgres driver
	"github.com/ory/dockertest/v3"
)

var db *sqlx.DB

var log = logrus.New()

// TestMain need to configure postgres test container for integration tests
func TestMain(m *testing.M) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}
	resource, err := pool.Run("postgres", "9.6", []string{"POSTGRES_USER=soypita", "POSTGRES_PASSWORD=soypita", "POSTGRES_DB=calendar"})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	dsn := fmt.Sprintf("postgres://soypita:soypita@localhost:%s/%s?sslmode=disable",
		resource.GetPort("5432/tcp"), "calendar")

	if err = pool.Retry(func() error {
		db, err = sqlx.Open("postgres", dsn)
		if err != nil {
			return err
		}
		return db.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	// run migration
	if err := goose.Run("up", db.DB, "../../../migrations"); err != nil {
		log.Fatalf("Could not run migration: %s", err)
	}

	code := m.Run()

	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	os.Exit(code)
}

func TestBasicPGRepository(t *testing.T) {
	t.Run(`should successfully create postgres repository`, func(t *testing.T) {
		repo := NewPGRepository(log, db)
		assert.NotNil(t, repo)
	})
	t.Run(`should successfully create event`, func(t *testing.T) {
		repo := NewPGRepository(log, db)
		resId, err := repo.CreateEvent(context.Background(), &models.Event{
			ID:           uuid.New(),
			Header:       "Test",
			Date:         time.Now(),
			Duration:     time.Duration(5) * time.Hour,
			Description:  "Test event",
			OwnerID:      uuid.New(),
			NotifyBefore: time.Duration(15) * time.Minute,
		})
		assert.NoError(t, err)
		assert.NotEqual(t, uuid.Nil, resId)
		cleanupRepository(t)
	})
	t.Run(`should get error when try to create event for busy date`, func(t *testing.T) {
		repo := NewPGRepository(log, db)
		insTime := time.Now()
		resId, err := repo.CreateEvent(context.Background(), &models.Event{
			ID:           uuid.New(),
			Header:       "Test",
			Date:         insTime,
			Duration:     time.Duration(5) * time.Hour,
			Description:  "Test event",
			OwnerID:      uuid.New(),
			NotifyBefore: time.Duration(15) * time.Minute,
		})
		assert.NoError(t, err)
		assert.NotEqual(t, uuid.Nil, resId)
		_, err = repo.CreateEvent(context.Background(), &models.Event{
			ID:           uuid.New(),
			Header:       "Test",
			Date:         insTime,
			Duration:     time.Duration(5) * time.Hour,
			Description:  "Test event",
			OwnerID:      uuid.New(),
			NotifyBefore: time.Duration(15) * time.Minute,
		})
		assert.Error(t, err)
		assert.Equal(t, repository.ErrDateBusy{}.Error(), err.Error())
		cleanupRepository(t)
	})
	t.Run(`should successfully create concurrently`, func(t *testing.T) {
		repo := NewPGRepository(log, db)
		assert.NotNil(t, repo)
		wg := sync.WaitGroup{}
		wg.Add(2)
		go func() {
			defer wg.Done()
			for i := 0; i < 100; i++ {
				resId, err := repo.CreateEvent(context.Background(), &models.Event{
					ID:           uuid.New(),
					Header:       fmt.Sprintf("Test %d from 2 goroutine", i),
					Date:         time.Now().Add(time.Duration(rand.Int()) * time.Minute),
					Duration:     time.Duration(5) * time.Hour,
					Description:  "Test event",
					OwnerID:      uuid.New(),
					NotifyBefore: time.Duration(15) * time.Minute,
				})
				assert.NoError(t, err)
				assert.NotEqual(t, uuid.Nil, resId)
			}
		}()
		go func() {
			defer wg.Done()
			for i := 0; i < 100; i++ {
				resId, err := repo.CreateEvent(context.Background(), &models.Event{
					ID:           uuid.New(),
					Header:       fmt.Sprintf("Test %d from 2 goroutine", i),
					Date:         time.Now().Add(time.Duration(rand.Int()) * time.Minute),
					Duration:     time.Duration(5) * time.Hour,
					Description:  "Test event",
					OwnerID:      uuid.New(),
					NotifyBefore: time.Duration(15) * time.Minute,
				})
				assert.NoError(t, err)
				assert.NotEqual(t, uuid.Nil, resId)
			}
		}()
		wg.Wait()
		events, err := repo.GetAllEvents(context.Background())
		assert.NoError(t, err)
		assert.Equal(t, 200, len(events))
		cleanupRepository(t)
	})
	t.Run(`should successfully update event in repository`, func(t *testing.T) {
		repo := NewPGRepository(log, db)
		resId, err := repo.CreateEvent(context.Background(), &models.Event{
			ID:           uuid.New(),
			Header:       "Test",
			Date:         time.Now(),
			Duration:     time.Duration(5) * time.Hour,
			Description:  "Test event",
			OwnerID:      uuid.UUID{},
			NotifyBefore: time.Duration(15) * time.Minute,
		})
		assert.NoError(t, err)
		assert.NotEqual(t, uuid.Nil, resId)
		updateEvent := &models.Event{
			ID:           resId,
			Header:       "Update Test",
			Date:         time.Time{},
			Duration:     0,
			Description:  "Updating event test",
			OwnerID:      uuid.UUID{},
			NotifyBefore: 0,
		}
		err = repo.UpdateEvent(context.Background(), resId, updateEvent)
		assert.NoError(t, err)
		resUpdate, err := repo.GetEventByID(context.Background(), resId)
		assert.NoError(t, err)
		assert.NotNil(t, resUpdate)
		assert.True(t, updateEvent.Date.Equal(resUpdate.Date))
		assert.Equal(t, updateEvent.Header, resUpdate.Header)
		assert.Equal(t, updateEvent.Duration, resUpdate.Duration)
		assert.Equal(t, updateEvent.Description, resUpdate.Description)
		assert.Equal(t, updateEvent.NotifyBefore, resUpdate.NotifyBefore)
		assert.Equal(t, updateEvent.OwnerID, resUpdate.OwnerID)
		cleanupRepository(t)
	})
	t.Run(`should get error when try to update event to busy date`, func(t *testing.T) {
		repo := NewPGRepository(log, db)
		busyDate := time.Now()
		resId, err := repo.CreateEvent(context.Background(), &models.Event{
			ID:           uuid.New(),
			Header:       "Test",
			Date:         busyDate,
			Duration:     time.Duration(5) * time.Hour,
			Description:  "Test event",
			OwnerID:      uuid.UUID{},
			NotifyBefore: time.Duration(15) * time.Minute,
		})
		assert.NoError(t, err)
		assert.NotEqual(t, uuid.Nil, resId)
		updateEvent := &models.Event{
			ID:           resId,
			Header:       "Update Test",
			Date:         busyDate,
			Duration:     0,
			Description:  "Updating event test",
			OwnerID:      uuid.UUID{},
			NotifyBefore: 0,
		}
		err = repo.UpdateEvent(context.Background(), resId, updateEvent)
		assert.Error(t, err, repository.ErrDateBusy{})
		assert.Equal(t, repository.ErrDateBusy{}.Error(), err.Error())
		cleanupRepository(t)
	})
	t.Run(`should successfully find event by id`, func(t *testing.T) {
		repo := NewPGRepository(log, db)
		insTime := time.Now()
		resId, err := repo.CreateEvent(context.Background(), &models.Event{
			ID:           uuid.New(),
			Header:       "Test",
			Date:         insTime,
			Duration:     time.Duration(5) * time.Hour,
			Description:  "Test event",
			OwnerID:      uuid.New(),
			NotifyBefore: time.Duration(15) * time.Minute,
		})
		assert.NoError(t, err)
		assert.NotEqual(t, uuid.Nil, resId)
		result, err := repo.GetEventByID(context.Background(), resId)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		cleanupRepository(t)
	})
	t.Run(`should return error when events not found`, func(t *testing.T) {
		repo := NewPGRepository(log, db)
		result, err := repo.GetEventByID(context.Background(), uuid.New())
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, repository.ErrEventNotFound{}.Error(), err.Error())
	})
	t.Run(`should successfully delete event from repo`, func(t *testing.T) {
		repo := NewPGRepository(log, db)
		resId, err := repo.CreateEvent(context.Background(), &models.Event{
			ID:           uuid.New(),
			Header:       "Test",
			Date:         time.Now(),
			Duration:     time.Duration(5) * time.Hour,
			Description:  "Test event",
			OwnerID:      uuid.UUID{},
			NotifyBefore: time.Duration(15) * time.Minute,
		})
		assert.NoError(t, err)
		assert.NotEqual(t, uuid.Nil, resId)
		err = repo.DeleteEvent(context.Background(), resId)
		assert.NoError(t, err)
		events, err := repo.GetAllEvents(context.Background())
		assert.NoError(t, err)
		assert.Empty(t, events)
	})
	t.Run(`should find all events for specific day`, func(t *testing.T) {
		repo := NewPGRepository(log, db)
		assert.NotNil(t, repo)
		firstTime, err := time.Parse("2006-01-02T15:04:05-0700", "2020-06-06T20:00:00-0300")
		if err != nil {
			panic(err)
		}
		secondTime, err := time.Parse("2006-01-02T15:04:05-0700", "2020-06-06T15:00:00-0300")
		if err != nil {
			panic(err)
		}
		thirdTime, err := time.Parse("2006-01-02T15:04:05-0700", "2020-05-06T15:00:00-0300")
		if err != nil {
			panic(err)
		}
		startDayTime, err := time.Parse("2006-01-02T15:04:05-0700", "2020-06-06T00:00:00-0300")
		if err != nil {
			panic(err)
		}
		firstEvent := &models.Event{
			ID:           uuid.New(),
			Header:       "Test",
			Date:         firstTime,
			Duration:     time.Duration(5) * time.Hour,
			Description:  "Test event",
			OwnerID:      uuid.UUID{},
			NotifyBefore: time.Duration(15) * time.Minute,
		}
		secondEvent := &models.Event{
			ID:           uuid.New(),
			Header:       "Test",
			Date:         secondTime,
			Duration:     time.Duration(5) * time.Hour,
			Description:  "Test event 2",
			OwnerID:      uuid.UUID{},
			NotifyBefore: time.Duration(15) * time.Minute,
		}
		thirdEvent := &models.Event{
			ID:           uuid.New(),
			Header:       "Test",
			Date:         thirdTime,
			Duration:     time.Duration(5) * time.Hour,
			Description:  "Test event 3",
			OwnerID:      uuid.UUID{},
			NotifyBefore: time.Duration(15) * time.Minute,
		}
		_, err = repo.CreateEvent(context.Background(), firstEvent)
		assert.NoError(t, err)
		_, err = repo.CreateEvent(context.Background(), secondEvent)
		assert.NoError(t, err)
		_, err = repo.CreateEvent(context.Background(), thirdEvent)
		assert.NoError(t, err)

		results, err := repo.FindDayEvents(context.Background(), startDayTime)
		assert.NoError(t, err)
		assert.NotEmpty(t, results)
		assert.Equal(t, 2, len(results))
		cleanupRepository(t)
	})
	t.Run(`should find event with date equal start day`, func(t *testing.T) {
		repo := NewPGRepository(log, db)
		assert.NotNil(t, repo)
		firstTime, err := time.Parse("2006-01-02T15:04:05-0700", "2020-06-06T20:00:00-0300")
		if err != nil {
			panic(err)
		}
		firstEvent := &models.Event{
			ID:           uuid.New(),
			Header:       "Test",
			Date:         firstTime,
			Duration:     time.Duration(5) * time.Hour,
			Description:  "Test event",
			OwnerID:      uuid.UUID{},
			NotifyBefore: time.Duration(15) * time.Minute,
		}
		_, err = repo.CreateEvent(context.Background(), firstEvent)
		assert.NoError(t, err)

		results, err := repo.FindDayEvents(context.Background(), firstTime)
		assert.NoError(t, err)
		assert.NotEmpty(t, results)
		assert.Equal(t, 1, len(results))
		cleanupRepository(t)
	})
	t.Run(`should find all events for specific week`, func(t *testing.T) {
		repo := NewPGRepository(log, db)
		assert.NotNil(t, repo)
		firstTime, err := time.Parse("2006-01-02T15:04:05-0700", "2020-06-06T20:00:00-0300")
		if err != nil {
			panic(err)
		}
		secondTime, err := time.Parse("2006-01-02T15:04:05-0700", "2020-06-08T15:00:00-0300")
		if err != nil {
			panic(err)
		}
		thirdTime, err := time.Parse("2006-01-02T15:04:05-0700", "2020-05-06T15:00:00-0300")
		if err != nil {
			panic(err)
		}
		startWeekTime, err := time.Parse("2006-01-02T15:04:05-0700", "2020-06-05T00:00:00-0300")
		if err != nil {
			panic(err)
		}
		firstEvent := &models.Event{
			ID:           uuid.New(),
			Header:       "Test",
			Date:         firstTime,
			Duration:     time.Duration(5) * time.Hour,
			Description:  "Test event",
			OwnerID:      uuid.UUID{},
			NotifyBefore: time.Duration(15) * time.Minute,
		}
		secondEvent := &models.Event{
			ID:           uuid.New(),
			Header:       "Test",
			Date:         secondTime,
			Duration:     time.Duration(5) * time.Hour,
			Description:  "Test event 2",
			OwnerID:      uuid.UUID{},
			NotifyBefore: time.Duration(15) * time.Minute,
		}
		thirdEvent := &models.Event{
			ID:           uuid.New(),
			Header:       "Test",
			Date:         thirdTime,
			Duration:     time.Duration(5) * time.Hour,
			Description:  "Test event 3",
			OwnerID:      uuid.UUID{},
			NotifyBefore: time.Duration(15) * time.Minute,
		}
		_, err = repo.CreateEvent(context.Background(), firstEvent)
		assert.NoError(t, err)
		_, err = repo.CreateEvent(context.Background(), secondEvent)
		assert.NoError(t, err)
		_, err = repo.CreateEvent(context.Background(), thirdEvent)
		assert.NoError(t, err)

		results, err := repo.FindWeekEvents(context.Background(), startWeekTime)
		assert.NoError(t, err)
		assert.NotEmpty(t, results)
		assert.Equal(t, 2, len(results))
		cleanupRepository(t)
	})
	t.Run(`should find event with date equal start day week`, func(t *testing.T) {
		repo := NewPGRepository(log, db)
		assert.NotNil(t, repo)
		firstTime, err := time.Parse("2006-01-02T15:04:05-0700", "2020-06-06T20:00:00-0300")
		if err != nil {
			panic(err)
		}
		firstEvent := &models.Event{
			ID:           uuid.New(),
			Header:       "Test",
			Date:         firstTime,
			Duration:     time.Duration(5) * time.Hour,
			Description:  "Test event",
			OwnerID:      uuid.UUID{},
			NotifyBefore: time.Duration(15) * time.Minute,
		}
		_, err = repo.CreateEvent(context.Background(), firstEvent)
		assert.NoError(t, err)

		results, err := repo.FindWeekEvents(context.Background(), firstTime)
		assert.NoError(t, err)
		assert.NotEmpty(t, results)
		assert.Equal(t, 1, len(results))
		cleanupRepository(t)
	})
	t.Run(`should find all events for specific month`, func(t *testing.T) {
		repo := NewPGRepository(log, db)
		assert.NotNil(t, repo)
		firstTime, err := time.Parse("2006-01-02T15:04:05-0700", "2020-05-29T20:00:00-0300")
		if err != nil {
			panic(err)
		}
		secondTime, err := time.Parse("2006-01-02T15:04:05-0700", "2020-06-03T15:00:00-0300")
		if err != nil {
			panic(err)
		}
		thirdTime, err := time.Parse("2006-01-02T15:04:05-0700", "2020-07-06T15:00:00-0300")
		if err != nil {
			panic(err)
		}
		startMonthTime, err := time.Parse("2006-01-02T15:04:05-0700", "2020-05-06T00:00:00-0300")
		if err != nil {
			panic(err)
		}
		firstEvent := &models.Event{
			ID:           uuid.New(),
			Header:       "Test",
			Date:         firstTime,
			Duration:     time.Duration(5) * time.Hour,
			Description:  "Test event",
			OwnerID:      uuid.UUID{},
			NotifyBefore: time.Duration(15) * time.Minute,
		}
		secondEvent := &models.Event{
			ID:           uuid.New(),
			Header:       "Test",
			Date:         secondTime,
			Duration:     time.Duration(5) * time.Hour,
			Description:  "Test event 2",
			OwnerID:      uuid.UUID{},
			NotifyBefore: time.Duration(15) * time.Minute,
		}
		thirdEvent := &models.Event{
			ID:           uuid.New(),
			Header:       "Test",
			Date:         thirdTime,
			Duration:     time.Duration(5) * time.Hour,
			Description:  "Test event 3",
			OwnerID:      uuid.UUID{},
			NotifyBefore: time.Duration(15) * time.Minute,
		}
		_, err = repo.CreateEvent(context.Background(), firstEvent)
		assert.NoError(t, err)
		_, err = repo.CreateEvent(context.Background(), secondEvent)
		assert.NoError(t, err)
		_, err = repo.CreateEvent(context.Background(), thirdEvent)
		assert.NoError(t, err)

		results, err := repo.FindMonthEvents(context.Background(), startMonthTime)
		assert.NoError(t, err)
		assert.NotEmpty(t, results)
		assert.Equal(t, 2, len(results))
		cleanupRepository(t)
	})
	t.Run(`should find event with date equal start day month`, func(t *testing.T) {
		repo := NewPGRepository(log, db)
		assert.NotNil(t, repo)
		firstTime, err := time.Parse("2006-01-02T15:04:05-0700", "2020-06-06T20:00:00-0300")
		if err != nil {
			panic(err)
		}
		firstEvent := &models.Event{
			ID:           uuid.New(),
			Header:       "Test",
			Date:         firstTime,
			Duration:     time.Duration(5) * time.Hour,
			Description:  "Test event",
			OwnerID:      uuid.UUID{},
			NotifyBefore: time.Duration(15) * time.Minute,
		}
		_, err = repo.CreateEvent(context.Background(), firstEvent)
		assert.NoError(t, err)

		results, err := repo.FindMonthEvents(context.Background(), firstTime)
		assert.NoError(t, err)
		assert.NotEmpty(t, results)
		assert.Equal(t, 1, len(results))
		cleanupRepository(t)
	})
}

func cleanupRepository(t *testing.T) {
	_, err := db.ExecContext(context.Background(), "DELETE FROM events")
	if err != nil {
		t.Error("error resetting:", err)
	}
}
