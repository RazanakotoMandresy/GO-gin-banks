package epargne

import (
	"net/http"

	"github.com/RazanakotoMandresy/go-gin-banks/pkg/middleware"
	"github.com/gin-gonic/gin"
)

// only epargne non economie (false) and with
func (h Handler) GetMoneyEpargne(ctx *gin.Context) {
	epargneUUID := ctx.Param("epargneUUID")
	userConnectedUUID, err := middleware.ExtractTokenUUID(ctx)
	if err != nil {
		return
	}
	// user to connected and to update
	usr := middleware.User{UuidToFind: userConnectedUUID, Db: h.DB}
	user, err := usr.User()
	if err != nil {
		return
	}
	singleEpargne := getSingleEparne{userUUID: user.UUID, epargneUUID: epargneUUID, h: h}
	epargne, err := singleEpargne.singleEconomie()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": "epargne already deleted or doesn't exist"})
		return
	}
	user.Moneys = user.Moneys + epargne.Value
	if res := h.DB.Save(&user); res.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": "error when saving user" + res.Error.Error()})
		return
	}

	if res := h.DB.Delete(&epargne); res.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": "error when deleting epargnes" + res.Error.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"res": epargne, "user": user})
}
