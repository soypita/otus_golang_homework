package models

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	ID           uuid.UUID     `json:"-"`
	Header       string        `json:"header"`
	Date         time.Time     `json:"date"`
	Duration     time.Duration `json:"duration"`
	Description  string        `json:"description"`
	OwnerID      uuid.UUID     `json:"owner_id"`
	NotifyBefore time.Duration `json:"notify_before"`
}
