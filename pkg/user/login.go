package user

import (
	"net/http"
	"time"

	"github.com/RazanakotoMandresy/bank-app-aout/backend/pkg/common/models"
	"github.com/RazanakotoMandresy/bank-app-aout/backend/pkg/middleware"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type LoginRequest struct {
	Email    string `json:"Email"`
	Password string `json:"password"`
}

func (h handler) Login(ctx *gin.Context) {
	body := new(LoginRequest)
	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	var users models.User
	email := models.User{Email: body.Email}

	GetPasswordHashed := h.DB.First(&users, email)
	if GetPasswordHashed.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, GetPasswordHashed.Error.Error())
		return
	}
	err := middleware.IsTruePassword(users.Password, body.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  users.ID,
		"uuid": users.UUID,
		"exp":  time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	tokenString, _ := middleware.TokenManage(token, ctx)
	ctx.JSON(http.StatusCreated, gin.H{"token": tokenString, "AppUserName": users.AppUserName})
}
