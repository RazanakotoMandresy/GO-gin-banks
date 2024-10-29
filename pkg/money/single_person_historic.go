package money

import (
	"net/http"

	"github.com/RazanakotoMandresy/go-gin-banks/pkg/common/models"
	"github.com/RazanakotoMandresy/go-gin-banks/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func (h handler) SingleHistoricTransaction(ctx *gin.Context) {
	body := new(historicRequest)
	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	uuidSended := ctx.Param("uuid")
	var transactionHistorics []models.Money
	uuid, err := middleware.ExtractTokenUUID(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	if body.Order == "" {
		// by default created_at
		body.Order = "created_at"
	}
	if body.Limit == 0 {
		body.Limit = 10
	}
	res := h.DB.Order(body.Order+" DESC").Limit(body.Limit).Find(&transactionHistorics, "send_by = ? AND sent_to = ?", uuid, uuidSended)
	if res.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, res.Error.Error())
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"historics": transactionHistorics})
}
