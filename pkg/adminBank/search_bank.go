package adminbank

import (
	"fmt"
	"net/http"

	"github.com/RazanakotoMandresy/bank-app-aout/backend/pkg/common/models"
	"github.com/RazanakotoMandresy/bank-app-aout/backend/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func (h handler) SearchBanks(ctx *gin.Context) {
	uuidAdmin, err := middleware.ExtractTokenUUID(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"err": err.Error()})
		return
	}
	// place de la banque pour le rechercher
	banquePlace := ctx.Query("place")
	var bank = models.Bank{}
	if result := h.DB.Where("maintenned_by = ? AND lieux = ?", uuidAdmin, banquePlace).First(&bank); result.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"err": fmt.Sprintf("la place %v n'existe pas ou bien ce n'es pas vous qui l'aver cree", banquePlace)})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"res": &bank})
}
