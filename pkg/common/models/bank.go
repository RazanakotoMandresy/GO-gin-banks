package models

import (
	"time"

	"gorm.io/gorm"
)

// models bank
type Bank struct {
	ID           uint32 `gorm:"id;primaryKey"`
	Created_at   time.Time
	Updated_at   time.Time
	Deleted_at   gorm.DeletedAt
	Money        int32
	Lieux        string `gorm:"unique"`
	MaintennedBy string
}
