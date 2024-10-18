package adminbank

import (
	"fmt"

	"github.com/RazanakotoMandresy/bank-app-aout/backend/pkg/common/models"
)

func (h handler) GetAdminUUID(adminUUID string) (*models.Admin, error) {
	var admin models.Admin
	result := h.DB.First(&admin, "uuid = ?", adminUUID)
	if result.Error != nil {
		err := fmt.Errorf("utilisateur avec l'id %v n'est pas dans %v ou bien vous n'etes pas admin", adminUUID, admin)
		return nil, err
	}
	return &admin, nil
}
