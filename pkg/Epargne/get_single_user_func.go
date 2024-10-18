package epargne
import (
	"fmt"

	"github.com/RazanakotoMandresy/bank-app-aout/backend/pkg/common/models"
)

func (h handler) GetUserSingleUserFunc(uuidToFind string) (*models.User, error) {
	var user models.User
	result := h.DB.First(&user, "uuid = ?", uuidToFind)
	if result.Error != nil {
		err := fmt.Errorf("user not in our database err : %v", result.Error)
		return nil, err
	}
	// if success user found user w t uuidToFins
	return &user, nil
}
