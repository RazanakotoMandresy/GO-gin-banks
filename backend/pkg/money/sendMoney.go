package money

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/RazanakotoMandresy/bank-app-aout/backend/pkg/common/models"
	"github.com/RazanakotoMandresy/bank-app-aout/backend/pkg/middleware"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lib/pq"
	// "github.com/google/uuid"
)

// ny uuid anaty params de ny uuid an'i envoyeur
type sendMoneyRequest struct {
	Value int32 `json:"value"`
}

func (h handler) SendMoney(ctx *gin.Context) {
	// extracttion de l'uuid depuis le bearer
	uuidConnectedStr, err := middleware.ExtractTokenUUID(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	// uuid du recepteur params dans l'url
	uuidRecepteur := ctx.Param("uuid")
	body := new(sendMoneyRequest)

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err})
		return
	}
	value := body.Value
	// code si l'on veux envoyer une somme inferieur a 1
	if value < 1 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": "on ne peut pas envoyer une somme aussi minime"})
		return
	}
	userConnected, _ := h.GetUserByuuid(uuidConnectedStr)
	// maka anle userRecepteur tokony par uuid na par AppUserName
	userRecepteur, err := h.GetUserByuuid(uuidRecepteur)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	// check si l'envoyeur essayent d'envoyer plus d'argent que ce qu'il en a
	if int(value) > userConnected.Moneys {
		err := fmt.Errorf("impossible d'envoyer votre argent %v l'argent que vous voulez envoyer est %v", userConnected.Moneys, value)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	// 	message si tous se passe bien
	// message := fmt.Sprintf("%v a envoye un argent d'un montant de %v a %v", userConnected.AppUserName, value, userRecepteur.AppUserName)
	userConnected.Moneys = userConnected.Moneys - int(value)
	userRecepteur.Moneys = (userRecepteur.Moneys + int(value))
	// save les money ao anaty user satria ireo tables fatsy uuid simplement ril
	h.DB.Save(userRecepteur)
	h.DB.Save(userConnected)
	// la models ho creena
	moneyTransaction, err := h.dbManipulationSendMoney(uuidConnectedStr, uuidRecepteur, body)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, &moneyTransaction)
}

func (h handler) dbManipulationSendMoney(uuidConnectedStr, uuidRecepteur string, body *sendMoneyRequest) (*models.Money, error) {
	var moneyTransaction models.Money
	res := h.DB.Where("send_by = ? AND sent_to = ?", uuidConnectedStr, uuidRecepteur).Find(&moneyTransaction)
	resume := fmt.Sprintf("%v a envoyer la somme de %v a %v a l'instant%v ", uuidConnectedStr, body.Value, uuidRecepteur, time.Now())
	if res.Error != nil {
		fmt.Printf("transaction entre send_by %v et sent_to %v inexistante creation d'une nouvelle...", uuidConnectedStr, uuidRecepteur)
		moneyTransaction.ID = uuid.New()
		moneyTransaction.SendBy = uuidConnectedStr
		moneyTransaction.SentTo = uuidRecepteur
		moneyTransaction.Totals = body.Value
		moneyTransaction.MoneyTransite = append(moneyTransaction.MoneyTransite, body.Value)
		moneyTransaction.Resume = resume
		moneyTransaction.MoneyTransite = append(moneyTransaction.MoneyTransite, body.Value)
		moneyTransaction.TransResum = append(moneyTransaction.TransResum, resume)
		result := h.DB.Create(moneyTransaction)
		if result.Error != nil {
			return nil, fmt.Errorf("creationraw %v", result.Error)
		}
	}
	moneyTransaction.ID = uuid.New()
	moneyTransaction.ID = uuid.New()
	moneyTransaction.SendBy = uuidConnectedStr
	moneyTransaction.SentTo = uuidRecepteur
	moneyTransaction.Resume = resume
	moneyTransaction.TransResum = append(moneyTransaction.TransResum, resume)
	moneyTransaction.MoneyTransite = append(moneyTransaction.MoneyTransite, body.Value)
	moneyTransaction.MoneyTransite = append(moneyTransaction.MoneyTransite, body.Value)
	totals := getTotals(moneyTransaction.MoneyTransite)
	// totals logique
	moneyTransaction.Totals = totals
	h.DB.Save(&moneyTransaction)
	return &moneyTransaction, nil
}
func getTotals(money pq.Int32Array) int32 {
	 var totals int32
	for i := 0; i < len(money); i++ {
		totals += money[i]
	}
	return totals
}

// user req que se soit uuid na appUserName
func (h handler) GetUserByuuid(userReq string) (*models.User, error) {
	var users models.User
	result := h.DB.First(&users, "uuid = ?", userReq)
	if result.Error != nil {
		err := fmt.Errorf("utilisateur avec l'uuid %v n'est pas dans %v , le resultats est %v recherche si c'est un appUserName", userReq, users, result)
		fmt.Println(err)
		res := h.DB.First(&users, "app_user_name = ?", userReq)
		if res.Error != nil {
			return nil, errors.New("user pas dans uuid et AppUserName")
		}
	}
	return &users, nil
}
