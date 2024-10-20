package user

import (
	"fmt"

	"github.com/RazanakotoMandresy/go-gin-banks/pkg/common/models"
)

func (h handler) GetUserSingleUserFunc(uuidToFind string) (*models.User, error) {
	var user models.User
	res := h.DB.Where("uuid = ? OR app_user_name = ?", uuidToFind, uuidToFind).First(&user)
	if res.Error != nil {
		return nil, fmt.Errorf(" %v pas dans uuid et AppUserName", uuidToFind)
	}
	return &user, nil
}
