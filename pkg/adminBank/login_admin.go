package adminbank

import (
	"net/http"
	"time"

	"github.com/RazanakotoMandresy/bank-app-aout/backend/pkg/common/models"
	"github.com/RazanakotoMandresy/bank-app-aout/backend/pkg/middleware"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type BankLogRequest struct {
	Name      string `json:"name"`
	Passwords string `json:"passwords"`
}

func (h handler) LoginAdmin(ctx *gin.Context) {
	body := new(BankLogRequest)
	if err := ctx.Bind(&body); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	if body.Name == "" || body.Passwords == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": "tous les champ sont obligatoire"})
		return
	}
	admin := models.Admin{Name: body.Name}
	GetHashedAdminPassword := h.DB.First(&admin, admin)
	if GetHashedAdminPassword.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, GetHashedAdminPassword.Error.Error())
		return
	}
	err := middleware.IsTruePassword(admin.Passwords, body.Passwords)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  admin.ID,
		"uuid": admin.UUID,
		"exp":  time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	tokenString, err := middleware.TokenManage(token, ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"token": tokenString})
}
