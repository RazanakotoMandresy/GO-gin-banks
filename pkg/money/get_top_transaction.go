package money

import (
	"net/http"

	"github.com/RazanakotoMandresy/go-gin-banks/pkg/common/models"
	"github.com/RazanakotoMandresy/go-gin-banks/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func (h handler) GetTopTrans(ctx *gin.Context) {
	uuid, err := middleware.ExtractTokenUUID(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err})
		return
	}
	var money []models.Money
	result := h.DB.Find(&money, "send_by = ?", uuid)
	if result.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": result.Error})
		return
	}
	userTosendSlicesJsoned := []topTrans{}
	// img for ui
	for _, moneys := range money {
		userTosendSlicesJsoned = append(userTosendSlicesJsoned, topTrans{
			SentTo:     moneys.SentTo,
			Totals:     int(moneys.Totals),
			UserName:   moneys.SentToName,
			ImageSento: moneys.SendToImg,
			ImgSender:  moneys.SendByImg,
		})
	}
	if len(userTosendSlicesJsoned) == 0 {
		ctx.JSON(http.StatusOK, gin.H{"res": "you haven't send money to anyone yet"})
		return
	}
	ctx.JSON(http.StatusOK, userTosendSlicesJsoned)
}
