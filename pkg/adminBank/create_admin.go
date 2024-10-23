package adminbank

import (
	"net/http"
	"time"

	"github.com/RazanakotoMandresy/go-gin-banks/pkg/common/models"
	"github.com/RazanakotoMandresy/go-gin-banks/pkg/middleware"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func (h handler) RegisterAdmin(ctx *gin.Context) {
	body := BankAdminReq{}
	if err := ctx.Bind(&body); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err})
		return
	}
	requriredField := map[string]string{
		"name":      body.Name,
		"passwords": body.Passwords,
		"root":      body.RootPass,
	}
	if !middleware.ValidateRequiredFields(ctx, requriredField) {
		return
	}
	err := middleware.IsTruePassword("$2a$05$/8PnBDSt7ZAxkdtW6c7.vOOusUebMZzT8ZMF4PtPc.DkI09XBi0I2", body.RootPass)
	// doesn't work
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	passwordHashed := middleware.HashPassword(body.Passwords)
	admin := models.Admin{
		ID:         uuid.New().ID(),
		UUID:       uuid.New(),
		Created_at: time.Now(),
		Updated_at: time.Now(),
		Name:       body.Name,
		Passwords:  passwordHashed,
		Role:       "admin",
	}
	// create the adminss in the db
	if result := h.DB.Create(&admin); result.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": result.Error})
		return
	}
	// signing the claims'token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  admin.ID,
		"uuid": admin.UUID,
		"exp":  time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	tokenString, _ := middleware.TokenManage(token, ctx)

	ctx.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}
