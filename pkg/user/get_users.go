package user

import (
	"net/http"

	"github.com/RazanakotoMandresy/bank-app-aout/backend/pkg/common/models"
	"github.com/gin-gonic/gin"
)

type userRes struct {
	UUID        string
	AppUserName string
	Email       string
}

func (h handler) GetUsers(ctx *gin.Context) {
	var users []models.User
	var usersFilterd []userRes
	res := h.DB.Find(&users).Limit(20)
	if res.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, res.Error)
		return
	}
	for _, user := range users {
		usersFilterd = append(usersFilterd, userRes{AppUserName: user.AppUserName, UUID: user.UUID, Email: user.Email})
	}
	ctx.JSON(http.StatusOK, gin.H{"res": usersFilterd})
}
