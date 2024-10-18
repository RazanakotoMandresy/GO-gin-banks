package chatrealtimes

import (
	"net/http"

	"github.com/RazanakotoMandresy/bank-app-aout/backend/pkg/common/models"
	"github.com/RazanakotoMandresy/bank-app-aout/backend/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func (h handler) GetAllMessageSent(ctx *gin.Context) {
	uuidSender, err := middleware.ExtractTokenUUID(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"err": err.Error(),
		})
	}
	uuidSentTo := ctx.Param("uuid")
	// pernd la liste des tous les discusion ou c'es lui l'envoyeur et l'envoyer
	var chat []models.Chat
	result := h.DB.Where("send_by = ? AND sent_to = ?", uuidSender, uuidSentTo).Order("created_at ASC").Find(&chat)
	if result.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": result.Error.Error()})
		return
	}
	ctx.JSON(http.StatusOK, &chat)
}
