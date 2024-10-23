package models

import (
	"time"

	"github.com/google/uuid"
)

type Epargne struct {
	ID           uuid.UUID `gorm:"id;primaryKey"`
	Created_at   time.Time
	Updated_at   time.Time
	Name         string
	DayPerMounth uint
	Type         string
	Value        int32
	UserUUID     string
}
