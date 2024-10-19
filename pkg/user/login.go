package user

import (
	"net/http"
	"time"

	"github.com/RazanakotoMandresy/bank-app-aout/backend/pkg/common/models"
	"github.com/RazanakotoMandresy/bank-app-aout/backend/pkg/middleware"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func (h handler) Login(ctx *gin.Context) {
	// var users models.User
	body := new(loginRequest)
	if err := ctx.Bind(&body); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	if body.Email == "" || body.Password == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": "you need to complete all the inputs"})
		return
	}
	user := models.User{Email: body.Email}
	// only the userEmail is not empty
	if GetPasswordHashed := h.DB.First(&user, user); GetPasswordHashed.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, GetPasswordHashed.Error.Error())
		return
	}
	if err := middleware.IsTruePassword(user.Password, body.Password); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	tokenString, err := middleware.TokenManage(jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  user.ID,
		"uuid": user.UUID,
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
	}), ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"token": tokenString})
}
