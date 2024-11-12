package epargne

import (
	"net/http"

	"github.com/RazanakotoMandresy/go-gin-banks/pkg/common/models"
	"github.com/RazanakotoMandresy/go-gin-banks/pkg/middleware"
	"github.com/gin-gonic/gin"
)

type getEpargnes struct {
	userUUID string
	h        Handler
}

func (h Handler) GetAllMyEpargnes(ctx *gin.Context) {
	userConnected, err := middleware.ExtractTokenUUID(ctx)
	if err != nil {
		return
	}
	getEpargnes := getEpargnes{userUUID: userConnected, h: h}
	myEpargnesEconomies, err := getEpargnes.getAllMyEpargnesFuncEconomies()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err})
		return
	}
	myEpargnesNonEconomie, err := getEpargnes.getAllMyEpargnesFuncNotEconomie()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"epargnes_economies": myEpargnesEconomies, "simples_epargnes": myEpargnesNonEconomie})
}

func (g getEpargnes) getAllMyEpargnesFuncEconomies() ([]models.Epargne, error) {
	var epargnes []models.Epargne
	res := g.h.DB.Where("owner_uuid = ? AND is_economie = true", g.userUUID).Find(&epargnes)
	if res.Error != nil {
		return nil, res.Error
	}
	return epargnes, nil
}
func (g getEpargnes) getAllMyEpargnesFuncNotEconomie() ([]models.Epargne, error) {
	var epargnes []models.Epargne
	res := g.h.DB.Where("owner_uuid = ? AND is_economie = false", g.userUUID).Find(&epargnes)
	if res.Error != nil {
		return nil, res.Error
	}
	return epargnes, nil
}
