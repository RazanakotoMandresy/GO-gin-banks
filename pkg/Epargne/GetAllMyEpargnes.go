package epargne

import (
	"net/http"

	"github.com/RazanakotoMandresy/go-gin-banks/pkg/common/models"
	"github.com/RazanakotoMandresy/go-gin-banks/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func (h Handler) GetAllMyEpargnes(ctx *gin.Context) {
	userConnected, err := middleware.ExtractTokenUUID(ctx)
	if err != nil {
		return
	}
	myEpargnesEconomies, err := getAllMyEpargnesFuncEconomies(userConnected, h)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err})
		return
	}
	myEpargnesNonEconomie, err := getAllMyEpargnesFuncNotEconomie(userConnected, h)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"epargnes_economies": myEpargnesEconomies, "simples_epargnes": myEpargnesNonEconomie})
}

func getAllMyEpargnesFuncEconomies(userUUID string, h Handler) ([]models.Epargne, error) {
	var epargnes []models.Epargne
	res := h.DB.Where("owner_uuid = ? AND is_economie = true", userUUID).Find(&epargnes)
	if res.Error != nil {
		return nil, res.Error
	}
	return epargnes, nil
}
func getAllMyEpargnesFuncNotEconomie(userUUID string, h Handler) ([]models.Epargne, error) {
	var epargnes []models.Epargne
	res := h.DB.Where("owner_uuid = ? AND is_economie = false", userUUID).Find(&epargnes)
	if res.Error != nil {
		return nil, res.Error
	}
	return epargnes, nil
}
