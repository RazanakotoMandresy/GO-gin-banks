package models

import (
	"time"

	"gorm.io/gorm"
)

// models bank
type Bank struct {
	ID           uint32         `gorm:"id;primaryKey"`
	Created_at   time.Time      `json:"created_at"`
	Updated_at   time.Time      `json:"updated_at"`
	Deleted_at   gorm.DeletedAt `json:"deleted_at"`
	Money        int32          `json:"money"`
	Lieux        string         `gorm:"unique"`
	MaintennedBy string         `json:"maintenned_by"`
}
