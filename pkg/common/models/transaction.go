package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

// models money
type Money struct {
	ID            uuid.UUID     `gorm:"id;primarykey"`
	Totals        int32         `json:"totals"`
	Created_at    time.Time
	Updated_at    time.Time
	Deleted_at    gorm.DeletedAt
	SendBy        string `json:"sentBy"`
	SentTo        string `json:"sentTo"`
	SentToName    string `json:"sentToName"`
	Resume        string `json:"resume"`
	SendByImg     string `json:"sendByImg"`
	SendToImg     string `json:"SendToImg"`
}
