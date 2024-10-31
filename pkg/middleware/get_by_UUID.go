package middleware

import (
	"fmt"

	"github.com/RazanakotoMandresy/go-gin-banks/pkg/common/models"
	"gorm.io/gorm"
)

type User struct {
	userModel  *models.User
	Db         *gorm.DB
	UuidToFind string
}

func (u User) User() (*models.User, error) {
	uuidToFind := u.UuidToFind
	res := u.Db.Where("uuid = ? OR app_user_name = ?", uuidToFind, uuidToFind).First(&u.userModel)
	if res.Error != nil {
		return nil, fmt.Errorf("user with uuid or appName %v not found ", uuidToFind)
	}
	return u.userModel, nil
}

type Admin struct {
	AdminModel models.Admin
	Db         *gorm.DB
	UuidToFind string
}

func (a Admin) Admin() (*models.Admin, error) {
	result := a.Db.First(&a.AdminModel, "uuid = ?", a.UuidToFind)
	if result.Error != nil {
		err := fmt.Errorf("admin with the uuid : %s does't exist", a.UuidToFind)
		return nil, err
	}
	return &a.AdminModel, nil
}
