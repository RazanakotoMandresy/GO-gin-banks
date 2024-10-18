package user

import (
	"fmt"
	"net/http"

	"github.com/RazanakotoMandresy/bank-app-aout/backend/pkg/common/models"
	"github.com/RazanakotoMandresy/bank-app-aout/backend/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func (h handler) getUser(ctx *gin.Context) {
	userParam := ctx.Param("user")
	users, err := h.GetUserSingleUserFunc(userParam)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	uuid, err := middleware.ExtractTokenUUID(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	totals, err := h.getTotals(ctx, uuid, users.UUID)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"total":         totals,
		"uuid":          users.UUID,
		"image":         users.Image,
		"Email":         users.Email,
		"nameFirstName": fmt.Sprint(users.Name, " ", users.FirstName),
		"AppUserName":   users.AppUserName,
	})
}
func (h handler) getTotals(ctx *gin.Context, uuidSender, uuidRecepteur string) (int32, error) {
	var totals models.Money
	res := h.DB.Where("send_by = ? AND sent_to = ?", uuidSender, uuidRecepteur).First(&totals)
	if res.Error != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return 0, fmt.Errorf("vous n'avez pas encore transferer de l'argent a %v", uuidRecepteur)
	}
	return totals.Totals, nil
}
