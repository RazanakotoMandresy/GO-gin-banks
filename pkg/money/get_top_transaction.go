package money

import (
	"net/http"

	"github.com/RazanakotoMandresy/bank-app-aout/backend/pkg/common/models"
	"github.com/RazanakotoMandresy/bank-app-aout/backend/pkg/middleware"
	"github.com/gin-gonic/gin"
)

type topTrans struct {
	SentTo     string `json:"sentTo"`
	Totals     int    `json:"sommeTrans"`
	UserName   string `json:"userName"`
	ImageSento string `json:"SentToImg"`
	ImgSender  string `json:"SentByImg"`
}

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

	ctx.JSON(http.StatusOK, userTosendSlicesJsoned)
}
