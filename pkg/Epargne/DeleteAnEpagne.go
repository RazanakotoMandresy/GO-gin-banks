package epargne

import (
	"net/http"

	"github.com/RazanakotoMandresy/go-gin-banks/pkg/middleware"
	"github.com/gin-gonic/gin"
)

// only non economie

func (h Handler) DeleteEpargne(ctx *gin.Context) {
	epargneUUID := ctx.Param("epargneUUID")
	userConnectedUUID, err := middleware.ExtractTokenUUID(ctx)
	if err != nil {
		return
	}
	singleEpargne := getSingleEparne{userUUID: userConnectedUUID, epargneUUID: epargneUUID, h: h}
	epargne, err := singleEpargne.singleEpargneNonEconomie()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": "epargne already deleted or doesn't exist"})
		return
	}
	if res := h.DB.Delete(&epargne); res.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": res.Error.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"res": "deleted"})
}
