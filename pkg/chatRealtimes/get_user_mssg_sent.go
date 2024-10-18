package chatrealtimes

import (
	"fmt"
	"net/http"

	"github.com/RazanakotoMandresy/bank-app-aout/backend/pkg/common/models"
	"github.com/RazanakotoMandresy/bank-app-aout/backend/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func (h handler) GetUsrMsgSent(ctx *gin.Context) {
	uuid, err := middleware.ExtractTokenUUID(ctx)
	fmt.Println(uuid)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"err": err.Error()})
		return
	}
	var chat []models.Chat
	res := h.DB.Where("send_by = ?", uuid).Find(&chat)
	if res.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": res.Error.Error()})
		return
	}
	ctx.JSON(http.StatusOK, &chat)
}
