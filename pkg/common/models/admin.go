package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Admin struct {
	ID         uint32         `gorm:"id;unique"`
	UUID       uuid.UUID      `gorm:"uuid;primaryKey"`
	Created_at time.Time      `json:"created_by"`
	Updated_at time.Time      `json:"updated_at"`
	Deleted_at gorm.DeletedAt `json:"deleted_at"`
	Name       string         `gorm:"name;unique"`
	Passwords  string         `json:"passwords"`
	Role       string         `json:"role"`
	TotalSend  uint           `json:"total_send"`
}
