package calendar

import (
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/soypita/otus_golang_homework/hw12_13_14_15_calendar/internal/models"
	"github.com/soypita/otus_golang_homework/hw12_13_14_15_calendar/internal/repository/inmemory"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCalendarService(t *testing.T) {
	log := logrus.New()
	t.Run(`should successfully create event`, func(t *testing.T) {
		repo := inmemory.NewInMemRepository(log)
		calendar := NewCalendar(repo)
		resId, err := calendar.CreateEvent(&models.Event{
			Header:       "Test",
			Date:         time.Now(),
			Duration:     time.Duration(5) * time.Hour,
			Description:  "Test event",
			OwnerID:      uuid.UUID{},
			NotifyBefore: time.Duration(15) * time.Minute,
		})
		assert.NoError(t, err)
		assert.NotEqual(t, uuid.Nil, resId)
	})

	t.Run(`should successfully update event`, func(t *testing.T) {
		repo := inmemory.NewInMemRepository(log)
		calendar := NewCalendar(repo)
		resId, err := calendar.CreateEvent(&models.Event{
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
		err = calendar.UpdateEvent(resId, updateEvent)
		assert.NoError(t, err)
		resUpdate, err := calendar.GetEventByID(resId)
		assert.NoError(t, err)
		assert.Equal(t, updateEvent, resUpdate)
	})
	t.Run(`should successfully delete event from repo`, func(t *testing.T) {
		repo := inmemory.NewInMemRepository(log)
		calendar := NewCalendar(repo)
		resId, err := calendar.CreateEvent(&models.Event{
			Header:       "Test",
			Date:         time.Now(),
			Duration:     time.Duration(5) * time.Hour,
			Description:  "Test event",
			OwnerID:      uuid.UUID{},
			NotifyBefore: time.Duration(15) * time.Minute,
		})
		assert.NoError(t, err)
		assert.NotEqual(t, uuid.Nil, resId)
		err = calendar.DeleteEvent(resId)
		assert.NoError(t, err)
		events, err := calendar.GetAllEvents()
		assert.NoError(t, err)
		assert.Empty(t, events)
	})
	t.Run(`should find all events for specific day`, func(t *testing.T) {
		repo := inmemory.NewInMemRepository(log)
		calendar := NewCalendar(repo)
		firstTime, err := time.Parse("2006-01-02T15:04:05-0700", "2020-06-06T20:00:00-0300")
		if err != nil {
			panic(err)
		}
		secondTime, err := time.Parse("2006-01-02T15:04:05-0700", "2020-06-06T15:00:00-0300")
		if err != nil {
			panic(err)
		}
		thirdTime, err := time.Parse("2006-01-02T15:04:05-0700", "2020-05-06T15:10:00-0300")
		if err != nil {
			panic(err)
		}
		startDayTime, err := time.Parse("2006-01-02T15:04:05-0700", "2020-06-06T00:00:00-0300")
		if err != nil {
			panic(err)
		}
		firstEvent := &models.Event{
			Header:       "Test",
			Date:         firstTime,
			Duration:     time.Duration(5) * time.Hour,
			Description:  "Test event",
			OwnerID:      uuid.UUID{},
			NotifyBefore: time.Duration(15) * time.Minute,
		}
		secondEvent := &models.Event{
			Header:       "Test",
			Date:         secondTime,
			Duration:     time.Duration(5) * time.Hour,
			Description:  "Test event 2",
			OwnerID:      uuid.UUID{},
			NotifyBefore: time.Duration(15) * time.Minute,
		}
		thirdEvent := &models.Event{
			Header:       "Test",
			Date:         thirdTime,
			Duration:     time.Duration(5) * time.Hour,
			Description:  "Test event 3",
			OwnerID:      uuid.UUID{},
			NotifyBefore: time.Duration(15) * time.Minute,
		}
		_, err = calendar.CreateEvent(firstEvent)
		assert.NoError(t, err)
		_, err = calendar.CreateEvent(secondEvent)
		assert.NoError(t, err)
		_, err = calendar.CreateEvent(thirdEvent)
		assert.NoError(t, err)

		results, err := calendar.FindDayEvents(startDayTime)
		assert.NoError(t, err)
		assert.NotEmpty(t, results)
		assert.Equal(t, 2, len(results))
		assert.Contains(t, results, firstEvent)
		assert.Contains(t, results, secondEvent)
		assert.NotContains(t, results, thirdEvent)
	})
	t.Run(`should find all events for specific week`, func(t *testing.T) {
		repo := inmemory.NewInMemRepository(log)
		calendar := NewCalendar(repo)
		firstTime, err := time.Parse("2006-01-02T15:04:05-0700", "2020-06-07T20:00:00-0300")
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
			Header:       "Test",
			Date:         firstTime,
			Duration:     time.Duration(5) * time.Hour,
			Description:  "Test event",
			OwnerID:      uuid.UUID{},
			NotifyBefore: time.Duration(15) * time.Minute,
		}
		secondEvent := &models.Event{
			Header:       "Test",
			Date:         secondTime,
			Duration:     time.Duration(5) * time.Hour,
			Description:  "Test event 2",
			OwnerID:      uuid.UUID{},
			NotifyBefore: time.Duration(15) * time.Minute,
		}
		thirdEvent := &models.Event{
			Header:       "Test",
			Date:         thirdTime,
			Duration:     time.Duration(5) * time.Hour,
			Description:  "Test event 3",
			OwnerID:      uuid.UUID{},
			NotifyBefore: time.Duration(15) * time.Minute,
		}
		_, err = calendar.CreateEvent(firstEvent)
		assert.NoError(t, err)
		_, err = calendar.CreateEvent(secondEvent)
		assert.NoError(t, err)
		_, err = calendar.CreateEvent(thirdEvent)
		assert.NoError(t, err)

		results, err := calendar.FindWeekEvents(startWeekTime)
		assert.NoError(t, err)
		assert.NotEmpty(t, results)
		assert.Equal(t, 2, len(results))
		assert.Contains(t, results, firstEvent)
		assert.Contains(t, results, secondEvent)
		assert.NotContains(t, results, thirdEvent)
	})
	t.Run(`should find all events for specific month`, func(t *testing.T) {
		repo := inmemory.NewInMemRepository(log)
		calendar := NewCalendar(repo)
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
			Header:       "Test",
			Date:         firstTime,
			Duration:     time.Duration(5) * time.Hour,
			Description:  "Test event",
			OwnerID:      uuid.UUID{},
			NotifyBefore: time.Duration(15) * time.Minute,
		}
		secondEvent := &models.Event{
			Header:       "Test",
			Date:         secondTime,
			Duration:     time.Duration(5) * time.Hour,
			Description:  "Test event 2",
			OwnerID:      uuid.UUID{},
			NotifyBefore: time.Duration(15) * time.Minute,
		}
		thirdEvent := &models.Event{
			Header:       "Test",
			Date:         thirdTime,
			Duration:     time.Duration(5) * time.Hour,
			Description:  "Test event 3",
			OwnerID:      uuid.UUID{},
			NotifyBefore: time.Duration(15) * time.Minute,
		}
		_, err = calendar.CreateEvent(firstEvent)
		assert.NoError(t, err)
		_, err = calendar.CreateEvent(secondEvent)
		assert.NoError(t, err)
		_, err = calendar.CreateEvent(thirdEvent)
		assert.NoError(t, err)

		results, err := calendar.FindMonthEvents(startMonthTime)
		assert.NoError(t, err)
		assert.NotEmpty(t, results)
		assert.Equal(t, 2, len(results))
		assert.Contains(t, results, firstEvent)
		assert.Contains(t, results, secondEvent)
		assert.NotContains(t, results, thirdEvent)
	})
}
