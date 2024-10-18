package adminbank

import (
	"net/http"
	"time"

	"github.com/RazanakotoMandresy/bank-app-aout/backend/pkg/common/models"
	"github.com/RazanakotoMandresy/bank-app-aout/backend/pkg/middleware"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Code to create an admin
type BankAdminReq struct {
	ID         uint32 `gorm:"id;primaryKey"`
	Created_at time.Time
	Updated_at time.Time
	Deleted_at gorm.DeletedAt
	Name       string `json:"name"`
	Passwords  string `json:"passwords"`
	RootPass   string `json:"root"`
}

func (h handler) CreateAdminAccount(ctx *gin.Context) {
	body := BankAdminReq{}
	if err := ctx.Bind(&body); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err})
		return
	}
	if body.RootPass == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": "les mots de passes de super admin est obligatoire"})
		return
	}
	err := middleware.IsTruePassword("$2a$05$/8PnBDSt7ZAxkdtW6c7.vOOusUebMZzT8ZMF4PtPc.DkI09XBi0I2", body.RootPass)
	// doesn't work
	// err := middleware.IsTruePassword(os.Getenv("SECRET_ADMIN"), body.RootPass)

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
