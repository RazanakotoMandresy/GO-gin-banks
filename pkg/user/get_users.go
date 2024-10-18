package user

import (
	"net/http"

	"github.com/RazanakotoMandresy/bank-app-aout/backend/pkg/common/models"
	"github.com/gin-gonic/gin"
)

func (h handler) GetUsers(ctx *gin.Context) {
	var users []models.User
	res := h.DB.Find(&users)
	if res.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, res.Error)
		return
	}
	ctx.JSON(http.StatusOK, &users)
}
