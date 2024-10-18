package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Admin struct {
	ID         uint32    `gorm:"id;unique"`
	UUID       uuid.UUID `gorm:"uuid;primaryKey"`
	Created_at time.Time
	Updated_at time.Time
	Deleted_at gorm.DeletedAt
	Name       string `gorm:"name;unique"`
	Passwords  string `json:"passwords"`
	Role       string
	TotalSend  int
}
