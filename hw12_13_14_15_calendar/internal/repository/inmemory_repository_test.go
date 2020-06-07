package repository

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github/soypita/otus_golang_homework/hw12_13_14_15_calendar/internal/models"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestBasicInMemoryRepository(t *testing.T) {
	log := logrus.New()
	t.Run(`should successfully create in memory repository`, func(t *testing.T) {
		repo := NewInMemRepository(log)
		assert.NotNil(t, repo)
	})
	t.Run(`should successfully create event in repository`, func(t *testing.T) {
		repo := NewInMemRepository(log)
		assert.NotNil(t, repo)
		res, err := repo.CreateEvent(context.Background(), &models.Event{
			ID:           uuid.UUID{},
			Header:       "Test",
			Date:         time.Now(),
			Duration:     time.Duration(5) * time.Hour,
			Description:  "Test event",
			OwnerID:      uuid.UUID{},
			NotifyBefore: time.Duration(15) * time.Minute,
		})
		assert.NoError(t, err)
		assert.NotNil(t, res)
	})
	t.Run(`should successfully create concurrently`, func(t *testing.T) {
		repo := NewInMemRepository(log)
		assert.NotNil(t, repo)
		wg := sync.WaitGroup{}
		wg.Add(2)
		go func() {
			defer wg.Done()
			for i := 0; i < 100; i++ {
				res, err := repo.CreateEvent(context.Background(), &models.Event{
					ID:           uuid.New(),
					Header:       fmt.Sprintf("Test %d from 2 goroutine", i),
					Date:         time.Now().Add(time.Duration(rand.Int()) * time.Minute),
					Duration:     time.Duration(5) * time.Hour,
					Description:  "Test event",
					OwnerID:      uuid.New(),
					NotifyBefore: time.Duration(15) * time.Minute,
				})
				assert.NoError(t, err)
				assert.NotNil(t, res)
			}
		}()
		go func() {
			defer wg.Done()
			for i := 0; i < 100; i++ {
				res, err := repo.CreateEvent(context.Background(), &models.Event{
					ID:           uuid.New(),
					Header:       fmt.Sprintf("Test %d from 2 goroutine", i),
					Date:         time.Now().Add(time.Duration(rand.Int()) * time.Minute),
					Duration:     time.Duration(5) * time.Hour,
					Description:  "Test event",
					OwnerID:      uuid.New(),
					NotifyBefore: time.Duration(15) * time.Minute,
				})
				assert.NoError(t, err)
				assert.NotNil(t, res)
			}
		}()
		wg.Wait()
		events, err := repo.GetAllEvents(context.Background())
		assert.NoError(t, err)
		assert.Equal(t, 200, len(events))
	})
	t.Run(`should get error when try to add event to busy date`, func(t *testing.T) {
		repo := NewInMemRepository(log)
		assert.NotNil(t, repo)
		busyDate := time.Now()
		res, err := repo.CreateEvent(context.Background(), &models.Event{
			ID:           uuid.New(),
			Header:       "Test",
			Date:         busyDate,
			Duration:     time.Duration(5) * time.Hour,
			Description:  "Test event",
			OwnerID:      uuid.UUID{},
			NotifyBefore: time.Duration(15) * time.Minute,
		})
		assert.NoError(t, err)
		assert.NotNil(t, res)
		_, err = repo.CreateEvent(context.Background(), &models.Event{
			ID:           uuid.New(),
			Header:       "Test",
			Date:         busyDate,
			Duration:     time.Duration(5) * time.Hour,
			Description:  "Test event",
			OwnerID:      uuid.UUID{},
			NotifyBefore: time.Duration(15) * time.Minute,
		})
		assert.Error(t, err, ErrDateBusy{})
		assert.Equal(t, ErrDateBusy{}.Error(), err.Error())
	})
	t.Run(`should successfully update event in repository`, func(t *testing.T) {
		repo := NewInMemRepository(log)
		res, err := repo.CreateEvent(context.Background(), &models.Event{
			ID:           uuid.UUID{},
			Header:       "Test",
			Date:         time.Now(),
			Duration:     time.Duration(5) * time.Hour,
			Description:  "Test event",
			OwnerID:      uuid.UUID{},
			NotifyBefore: time.Duration(15) * time.Minute,
		})
		assert.NoError(t, err)
		assert.NotNil(t, res)
		updateEvent := &models.Event{
			ID:           res.ID,
			Header:       "Update Test",
			Date:         time.Time{},
			Duration:     0,
			Description:  "Updating event test",
			OwnerID:      uuid.UUID{},
			NotifyBefore: 0,
		}
		err = repo.UpdateEvent(context.Background(), res.ID, updateEvent)
		assert.NoError(t, err)
		resUpdate, err := repo.GetEventByID(context.Background(), res.ID)
		assert.NoError(t, err)
		assert.Equal(t, updateEvent, resUpdate)
	})
	t.Run(`should get error when try to update event to busy date`, func(t *testing.T) {
		repo := NewInMemRepository(log)
		busyDate := time.Now()
		res, err := repo.CreateEvent(context.Background(), &models.Event{
			ID:           uuid.UUID{},
			Header:       "Test",
			Date:         busyDate,
			Duration:     time.Duration(5) * time.Hour,
			Description:  "Test event",
			OwnerID:      uuid.UUID{},
			NotifyBefore: time.Duration(15) * time.Minute,
		})
		assert.NoError(t, err)
		assert.NotNil(t, res)
		updateEvent := &models.Event{
			ID:           res.ID,
			Header:       "Update Test",
			Date:         busyDate,
			Duration:     0,
			Description:  "Updating event test",
			OwnerID:      uuid.UUID{},
			NotifyBefore: 0,
		}
		err = repo.UpdateEvent(context.Background(), res.ID, updateEvent)
		assert.Error(t, err, ErrDateBusy{})
		assert.Equal(t, ErrDateBusy{}.Error(), err.Error())
	})
	t.Run(`should successfully delete event from repo`, func(t *testing.T) {
		repo := NewInMemRepository(log)
		res, err := repo.CreateEvent(context.Background(), &models.Event{
			ID:           uuid.UUID{},
			Header:       "Test",
			Date:         time.Now(),
			Duration:     time.Duration(5) * time.Hour,
			Description:  "Test event",
			OwnerID:      uuid.UUID{},
			NotifyBefore: time.Duration(15) * time.Minute,
		})
		assert.NoError(t, err)
		assert.NotNil(t, res)
		err = repo.DeleteEvent(context.Background(), res.ID)
		assert.NoError(t, err)
		events, err := repo.GetAllEvents(context.Background())
		assert.NoError(t, err)
		assert.Empty(t, events)
	})
	t.Run(`should find all events for specific day`, func(t *testing.T) {
		repo := NewInMemRepository(log)
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
		assert.Contains(t, results, firstEvent)
		assert.Contains(t, results, secondEvent)
		assert.NotContains(t, results, thirdEvent)
	})
	t.Run(`should find event with date equal start day`, func(t *testing.T) {
		repo := NewInMemRepository(log)
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
		assert.Contains(t, results, firstEvent)
	})
	t.Run(`should find all events for specific week`, func(t *testing.T) {
		repo := NewInMemRepository(log)
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
		assert.Contains(t, results, firstEvent)
		assert.Contains(t, results, secondEvent)
		assert.NotContains(t, results, thirdEvent)
	})
	t.Run(`should find event with date equal start day week`, func(t *testing.T) {
		repo := NewInMemRepository(log)
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
		assert.Contains(t, results, firstEvent)
	})
	t.Run(`should find all events for specific month`, func(t *testing.T) {
		repo := NewInMemRepository(log)
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
		assert.Contains(t, results, firstEvent)
		assert.Contains(t, results, secondEvent)
		assert.NotContains(t, results, thirdEvent)
	})
	t.Run(`should find event with date equal start day month`, func(t *testing.T) {
		repo := NewInMemRepository(log)
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
		assert.Contains(t, results, firstEvent)
	})
}
