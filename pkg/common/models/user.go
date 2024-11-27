package models

import (
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

// modles user
type User struct {
	Created_at  time.Time
	Updated_at  time.Time
	UUID        string `gorm:"uuid;primarykey"`
	Deleted_at  gorm.DeletedAt
	AppUserName string         `gorm:"unique"`
	Name        string         `json:"name"`
	Email       string         `gorm:"unique"`
	FirstName   string         `json:"firstName"`
	Moneys      uint           `json:"money"`
	Password    string         `json:"passwords"`
	BirthDate   string         `json:"dateNaissance"`
	Residance   string         `json:"residance"`
	Role        string         `json:"role"`
	Image       string         `json:"image"`
	BlockedAcc  pq.StringArray `gorm:"type:text[]"`
}
