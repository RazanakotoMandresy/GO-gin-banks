package adminbank

import (
	"time"

	"gorm.io/gorm"
)
type BankAdminReq struct {
	ID         uint32 `gorm:"id;primaryKey"`
	Created_at time.Time
	Updated_at time.Time
	Deleted_at gorm.DeletedAt
	Name       string `json:"name"`
	Passwords  string `json:"passwords"`
	RootPass   string `json:"root"`
}
type BankReq struct {
	Money    uint  `json:"money"`
	Lieux    string `json:"lieux"`
	Password string `json:"password"`
}
type BankLogRequest struct {
	Name      string `json:"name"`
	Passwords string `json:"passwords"`
}