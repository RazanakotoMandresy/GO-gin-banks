package epargne

import (
	"fmt"
	"net/http"
	"time"

	"github.com/RazanakotoMandresy/go-gin-banks/pkg/common/models"
	"github.com/RazanakotoMandresy/go-gin-banks/pkg/middleware"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreateEpargneRequest struct {
	Name    string `json:"name"`
	Type    string `json:"type"`
	Message string `json:"message"`
	// suppused to be the appUserName of the user sent to and then return it's uuid
	Sent_to    string `json:"sent_to"`
	Value      int32  `json:"value_epargne"`
	Date       uint   `json:"day_epargned"`
	AutoSend   bool   `json:"auto_send"`
	IsEconomie bool   `json:"is_economie"`
}

func (h Handler) CreateEpargne(ctx *gin.Context) {
	body := new(CreateEpargneRequest)
	if err := ctx.Bind(&body); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	userConnectedUUID, err := middleware.ExtractTokenUUID(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"err": err.Error()})
		return
	}
	user, err := middleware.User.User(middleware.User{UuidToFind: userConnectedUUID, Db: h.DB})
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"err": err.Error()})
		return
	}
	// logic stuff
	if body.Value > user.Moneys {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"err": fmt.Sprintf("vous ne pouvez pas epargner %v car l'argent sur votre compte est %v", body.Value, user.Moneys),
		})
		return
	}
	if body.Value == 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": "value are required"})
		return
	}
	requiredEpargne := map[string]string{
		"name":    body.Name,
		"type":    body.Type,
		"message": body.Message,
		"sent_to": body.Sent_to,
	}
	if !middleware.ValidateRequiredFields(ctx, requiredEpargne) {
		return
	}
	if body.IsEconomie {
		// handle logics economie if not economie default
		body.Sent_to = userConnectedUUID
		// economie always manuel get
		body.AutoSend = false
	}
	userToSend, err := middleware.User.User(middleware.User{UuidToFind: body.Sent_to, Db: h.DB})
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"err": err.Error()})
		return
	}
	epargne := models.Epargne{
		ID:           uuid.New(),
		Name:         body.Name,
		Value:        body.Value,
		DayPerMounth: body.Date,
		Type:         body.Type,
		OwnerUUID:    user.UUID,
		Message:      body.Message,
		Sent_to:      userToSend.UUID,
		// false non autosend
		AutoSend:   body.AutoSend,
		IsEconomie: body.IsEconomie,
		Created_at: time.Now(),
	}
	if res := h.DB.Create(&epargne); res.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": res.Error.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"epargne": &epargne})
}
