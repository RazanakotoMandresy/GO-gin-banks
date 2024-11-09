package models

import (
	"time"

	"github.com/google/uuid"
)

type Epargne struct {
	ID           uuid.UUID `gorm:"id;primaryKey"`
	Name         string
	Type         string
	OwnerUUID    string
	Message      string
	Sent_to      string
	DayPerMounth uint
	Value        int32
	AutoSend     bool
	IsEconomie   bool
	Updated_at   time.Time
	Deleted_at   time.Time
	Created_at   time.Time
}
type EpargneResume struct {
	ID            uuid.UUID `gorm:"id;primaryKey"`
	Type          string
	OwnerUUID     string
	ResumeMessage string
	Value         uint
	Created_at    time.Time
	Deleted_at    time.Time
}
