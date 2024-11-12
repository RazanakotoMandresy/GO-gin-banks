package epargne

import (
	"net/http"

	"github.com/RazanakotoMandresy/go-gin-banks/pkg/middleware"
	"github.com/gin-gonic/gin"
)


func (h Handler) GetAllMyEpargnes(ctx *gin.Context) {
	userConnected, err := middleware.ExtractTokenUUID(ctx)
	if err != nil {
		return
	}
	getEpargnes := getEpargnes{userUUID: userConnected, h: h}
	myEpargnesEconomies, err := getEpargnes.economieCase()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err})
		return
	}
	myEpargnesNonEconomie, err := getEpargnes.nonEconomieCase()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"epargnes_economies": myEpargnesEconomies, "simples_epargnes": myEpargnesNonEconomie})
}

