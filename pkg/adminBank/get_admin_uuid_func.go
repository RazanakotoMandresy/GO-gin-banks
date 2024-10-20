package adminbank

import (
	"fmt"

	"github.com/RazanakotoMandresy/go-gin-banks/pkg/common/models"
)

func (h handler) GetAdminUUID(adminUUID string) (*models.Admin, error) {
	var admin models.Admin
	result := h.DB.First(&admin, "uuid = ?", adminUUID)
	if result.Error != nil {
		err := fmt.Errorf("admin with the uuid : %s does't exist", adminUUID)
		return nil, err
	}
	return &admin, nil
}
