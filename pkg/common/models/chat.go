package models

import (
	"time"

	"github.com/google/uuid"
)

type Chat struct {
	ID uuid.UUID `gorm:"id;primaryKey"`
	Content string
	Created_at time.Time
	SendBy     string
	SentTo     string
}
