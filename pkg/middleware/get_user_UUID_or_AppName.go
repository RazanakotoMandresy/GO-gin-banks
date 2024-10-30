package middleware

import (
	"fmt"

	"github.com/RazanakotoMandresy/go-gin-banks/pkg/common/models"
	"gorm.io/gorm"
)

func GetUserUUID(db *gorm.DB, uuidToFind string) (*models.User, error) {
	var user models.User
	res := db.Where("uuid = ? OR app_user_name = ?", uuidToFind, uuidToFind).First(&user)
	if res.Error != nil {
		return nil, fmt.Errorf(" %v not found ", uuidToFind)
	}
	return &user, nil
}
