package adminbank

import (
	"net/http"
	"time"

	"github.com/RazanakotoMandresy/bank-app-aout/backend/pkg/common/models"
	"github.com/RazanakotoMandresy/bank-app-aout/backend/pkg/middleware"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func (h handler) LoginAdmin(ctx *gin.Context) {
	body := new(BankLogRequest)
	if err := ctx.Bind(&body); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	requiredField := map[string]string{
		"name":      body.Name,
		"passwords": body.Passwords,
	}
	if !middleware.ValidateRequiredFields(ctx, body, requiredField) {
		return
	}
	admin := models.Admin{Name: body.Name}
	// only the admin name is not empty
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
	tokenString, err := middleware.TokenManage(jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  admin.ID,
		"uuid": admin.UUID,
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
	}), ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"token": tokenString})
}
