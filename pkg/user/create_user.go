package user

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/RazanakotoMandresy/go-gin-banks/pkg/common/models"
	"github.com/RazanakotoMandresy/go-gin-banks/pkg/middleware"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func (h handler) CreateUser(ctx *gin.Context) {
	body := registerRequest{}
	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	requiredFields := map[string]string{
		"AppUserName": body.AppUserName,
		"BirthDate":   body.BirthDate,
		"Email":       body.Email,
		"FirstName":   body.FirstName,
		"Name":        body.Name,
		"Password":    body.Password,
	}
	if !middleware.ValidateRequiredFields(ctx, body, requiredFields) {
		return
	}
	passwordHashed := middleware.HashPassword(body.Password)
	user := models.User{
		AppUserName: body.AppUserName,
		Name:        body.Name,
		FirstName:   body.FirstName,
		Password:    passwordHashed,
		BirthDate:   body.BirthDate,
		Moneys:      0,
		UUID:        uuid.New().String(),
		Residance:   body.Residance,
		Email:       body.Email,
		Created_at:  time.Now(),
		Updated_at:  time.Now(),
		Role:        "user",
		Image:       "imgDef/defaultPP.jpg",
		BlockedAcc:  []string{},
	}

	if result := h.DB.Create(&user); result.Error != nil {
		strErr := result.Error.Error()
		// cannot use a real check cz the errors happen in differents languages
		if strings.ContainsAny(strErr, "23505") {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": fmt.Sprintf("cannot duplicate: -%v", strErr)})
			return
		}
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": result.Error.Error()})
		return
	}
	tokenString, _ := middleware.TokenManage(jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uuid": user.UUID,
		"role": user.Role,
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
	}), ctx)
	ctx.JSON(http.StatusCreated, gin.H{"token": tokenString})
}
