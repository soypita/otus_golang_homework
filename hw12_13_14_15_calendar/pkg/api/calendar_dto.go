package api

import (
	"time"

	"github.com/google/uuid"
)

type EventDTO struct {
	Header       string        `json:"header"`
	Date         time.Time     `json:"date"`
	Duration     time.Duration `json:"duration"`
	Description  string        `json:"description"`
	OwnerID      uuid.UUID     `json:"owner_id"`
	NotifyBefore time.Duration `json:"notify_before"`
}

type CreateEventRequest struct {
	Event *EventDTO `json:"event"`
}

type CreateEventResponse struct {
	ID string `json:"id"`
}

type EventUpdateRequest struct {
	Event *EventDTO `json:"event"`
}

type GetAllEventsResponse struct {
	Events []*EventDTO `json:"events"`
}

type FindEventByIDResponse struct {
	Event *EventDTO `json:"event"`
}

type FindDayEventsResponse struct {
	Events []*EventDTO `json:"events"`
}

type FindWeekEventsResponse struct {
	Events []*EventDTO `json:"events"`
}

type FindMonthEventsResponse struct {
	Events []*EventDTO `json:"events"`
}
