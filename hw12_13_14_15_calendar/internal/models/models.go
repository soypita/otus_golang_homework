package models

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	ID           uuid.UUID
	Header       string
	Date         time.Time
	Duration     time.Duration
	Description  string
	OwnerID      uuid.UUID
	NotifyBefore time.Duration
}
