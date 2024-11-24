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
	Money    int32  `json:"money"`
	Lieux    string `json:"place"`
	Password string `json:"passwords"`
}
type BankLogRequest struct {
	Name      string `json:"name"`
	Passwords string `json:"passwords"`
}