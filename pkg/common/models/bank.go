package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// models bank
type Bank struct {
	ID           uuid.UUID      `gorm:"id;primaryKey"`
	Created_at   time.Time      `json:"created_at"`
	Updated_at   time.Time      `json:"updated_at"`
	Deleted_at   gorm.DeletedAt `json:"deleted_at"`
	Money        uint           `json:"money"`
	Lieux        string         `gorm:"unique"`
	MaintennedBy string         `json:"maintenned_by"`
}
