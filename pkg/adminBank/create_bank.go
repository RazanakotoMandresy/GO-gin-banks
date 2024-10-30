package adminbank

import (
	"math/rand"
	"net/http"

	"github.com/RazanakotoMandresy/go-gin-banks/pkg/common/models"
	"github.com/RazanakotoMandresy/go-gin-banks/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func (h handler) CreateBank(ctx *gin.Context) {
	body := new(BankReq)
	// mamadika anle uuidAny ho string
	uuid, err := middleware.ExtractTokenUUID(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	admin, err := middleware.Admin.Admin(middleware.Admin{Db: h.DB, UuidToFind: uuid})
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}
	// check roles
	if admin.Role != "admin" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"err": "vous n'avez pas le role necessaire ",
		})
		return
	}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}
	if err := middleware.IsTruePassword(admin.Passwords, body.Password); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	bank := models.Bank{
		Money:        body.Money,
		Lieux:        body.Lieux,
		ID:           rand.Uint32(),
		MaintennedBy: admin.UUID.String(),
	}

	if result := h.DB.Create(&bank); result.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"err": result.Error,
		})
		return
	}

	admin.TotalSend = admin.TotalSend + int(body.Money)
	h.DB.Save(&admin)
	ctx.JSON(http.StatusOK, gin.H{
		"res": bank,
	})
}
