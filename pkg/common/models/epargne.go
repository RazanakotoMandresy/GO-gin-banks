package models

import (
	"time"

	"github.com/google/uuid"
)

type Epargne struct {
	ID         uuid.UUID `gorm:"id;primaryKey"`
	Created_at time.Time
	Updated_at time.Time
	Name       string
	// chaque jours du mois tonga de verser automatiquement
	DayPerMounth uint
	// Type genre type anle epargne asina 2 , versment auto dans un compte na vola angonona tsy afaka kitihina (genre boitendrakitra)
	Type     string
	Value    int32
	UserUUID string
}
