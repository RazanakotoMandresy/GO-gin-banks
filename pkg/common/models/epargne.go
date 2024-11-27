package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Epargne struct {
	ID           uuid.UUID      `gorm:"id;primaryKey" json:"id"`
	Name         string         `json:"name"`
	Type         string         `json:"type"`
	OwnerUUID    string         `json:"owner_uuid"`
	Message      string         `json:"message"`
	Sent_to      string         `json:"sent_to"`
	DayPerMounth int           `json:"day_per_month"`
	Value        uint          `json:"value"`
	AutoSend     bool           `json:"auto_send"`
	IsEconomie   bool           `json:"is_economie"`
	Deleted_at   gorm.DeletedAt `json:"deleted_at"`
	Created_at   time.Time      `json:"created_at"`
}
type EpargneResume struct {
	ID            uuid.UUID      `gorm:"id;primaryKey"`
	Type          string         `json:"type"`
	OwnerUUID     string         `json:"owner_uuid"`
	ResumeMessage string         `json:"resume_message"`
	Value         uint           `json:"value"`
	Created_at    time.Time      `json:"created_at"`
	Deleted_at    gorm.DeletedAt `json:"deleted_at"`
}
