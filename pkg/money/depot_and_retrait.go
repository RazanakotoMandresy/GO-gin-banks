package money

import (
	"fmt"
	"net/http"

	"github.com/RazanakotoMandresy/bank-app-aout/backend/pkg/common/models"
	"github.com/RazanakotoMandresy/bank-app-aout/backend/pkg/middleware"
	"github.com/gin-gonic/gin"
)

type DepoRetraiReq struct {
	Value     int32  `json:"value"`
	Lieux     string `json:"lieux"`
	Passwords string `json:"passwords"`
}

func (h handler) Depot(ctx *gin.Context) {
	usr, err := middleware.ExtractTokenUUID(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"err": err.Error()})
		return
	}
	body := new(DepoRetraiReq)
	userTosend, err := h.GetUserByuuid(usr)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	if body.Lieux == "" || body.Value == 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": "Lieux et value sont obligatoire"})
		return
	}
	moneyReq := body.Value
	var bank models.Bank
	result := h.DB.First(&bank, "lieux = ?", body.Lieux)
	if result.Error != nil {
		err := fmt.Sprintf("le lieux %v n'existe pas", body.Lieux)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"err": err,
		})
		return
	}
	// check si l'argent que l'user veut deposer est superieur a ce que la depotoire de la banque peut donner

	if bank.Money < moneyReq {
		err := fmt.Sprintf("erreur , l'argent que vous voulez depose est %v qui est largement superieur a la valeur du money virtuel dans ce bank %v", moneyReq, bank.Money)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err})
		return

	}

	userTosend.Moneys = (userTosend.Moneys + moneyReq)
	bank.Money = (bank.Money - moneyReq)

	h.DB.Save(bank)
	h.DB.Save(userTosend)

	msg := fmt.Sprintf("vous venez de faire un depot de %v votre solde est de : %v", moneyReq, userTosend.Moneys)
	ctx.JSON(http.StatusOK, gin.H{"msg": msg})

}
func (h handler) Retrait(ctx *gin.Context) {

	userTosendUUid, err := middleware.ExtractTokenUUID(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	body := new(DepoRetraiReq)

	userTosend, err := h.GetUserByuuid(userTosendUUid)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	if body.Lieux == "" || body.Value == 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": "Lieux et value sont obligatoire"})
		return
	}
	moneyReq := body.Value
	var bank models.Bank
	result := h.DB.First(&bank, "lieux = ?", body.Lieux)
	if result.Error != nil {
		err := fmt.Sprintf("le lieux %v n'existe pas", body.Lieux)

		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"err": err,
		})
		return
	}
	// check si l'argent que l'user veut deposer est superieur a ce que la depotoire de la banque peut donner
	if userTosend.Moneys < moneyReq {
		err := fmt.Sprintf("erreur , l'argent que vous voulez retirer est %v mais votre argents est de :  %v", moneyReq, userTosend.Moneys)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err})
		return

	}

	if err := middleware.IsTruePassword(userTosend.Password, body.Passwords); err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"err": err.Error()})
		return
	}
	userTosend.Moneys = (userTosend.Moneys - moneyReq)
	bank.Money = (bank.Money + moneyReq)

	h.DB.Save(bank)
	h.DB.Save(userTosend)

	msg := fmt.Sprintf("vous venez de faire un retrait de %v votre solde est de : %v", moneyReq, userTosend.Moneys)
	ctx.JSON(http.StatusOK, gin.H{"msg": msg})
}
