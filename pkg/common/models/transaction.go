package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

// models money
type Money struct {
	ID         uuid.UUID      `gorm:"id;primarykey" json:"uuid"`
	Totals     uint           `json:"totals"`
	Created_at time.Time      `json:"created_at"`
	Updated_at time.Time      `json:"updated_at"`
	Deleted_at gorm.DeletedAt `json:"deleted_at"`
	SendBy     string         `json:"sentBy"`
	SentTo     string         `json:"sentTo"`
	SentToName string         `json:"sentToName"`
	Resume     string         `json:"resume"`
	SendByImg  string         `json:"sendByImg"`
	SendToImg  string         `json:"SendToImg"`
}
