package epargne

import (
	"fmt"
	"net/http"

	"github.com/RazanakotoMandresy/go-gin-banks/pkg/common/models"
	"github.com/RazanakotoMandresy/go-gin-banks/pkg/middleware"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreateEpargneRequest struct {
	Name  string `json:"name"`
	Value int32  `json:"value_epargne"`
	Date  uint   `json:"day_epargned"`
	Type  string `json:"type"`
}

func (h handler) CreateEpargne(ctx *gin.Context) {
	body := new(CreateEpargneRequest)
	if err := ctx.Bind(&body); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err})
		return
	}
	userConnectedUUID, err := middleware.ExtractTokenUUID(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"err": err.Error()})
		return
	}
	// user, err := h.GetUserSingleUserFunc(userConnectedUUID)
	user, err := middleware.User.User(middleware.User{UuidToFind: userConnectedUUID, Db: h.DB})
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"err": err.Error()})
	}
	// logic stuff
	if body.Value > user.Moneys {
		err := fmt.Sprintf("vous ne pouvez pas epargner %v car l'argent sur votre compte est %v", body.Value, user.Moneys)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"err": err,
		})
		return
	}
	if body.Date == 0 || body.Value == 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": "date and value are required"})
		return
	}
	requiredEpargne := map[string]string{
		"name": body.Name,
		"type": body.Type,
	}
	if !middleware.ValidateRequiredFields(ctx, requiredEpargne) {
		return
	}
	epargne := models.Epargne{
		ID:           uuid.New(),
		Name:         body.Name,
		Value:        body.Value,
		DayPerMounth: body.Date,
		Type:         body.Type,
		OwnerUUID:    user.UUID,
	}
	h.DB.Create(&epargne)
	ctx.JSON(http.StatusOK, gin.H{"epargne": &epargne, "user": &user})
}
