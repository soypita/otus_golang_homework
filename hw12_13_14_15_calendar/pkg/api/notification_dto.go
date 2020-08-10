package api

import (
	"time"

	"github.com/google/uuid"
)

type NotificationDTO struct {
	EventID      uuid.UUID `json:"event_id"`
	EventHeader  string    `json:"event_header"`
	EventDate    time.Time `json:"event_date"`
	EventOwnerID uuid.UUID `json:"event_owner_id"`
}
