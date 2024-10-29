package money

import (
	"fmt"
	"net/http"
	"slices"
	"time"

	// "time"

	"github.com/RazanakotoMandresy/go-gin-banks/pkg/common/models"
	"github.com/RazanakotoMandresy/go-gin-banks/pkg/middleware"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

func (h handler) SendMoney(ctx *gin.Context) {
	// extracttion de l'uuid depuis le bearer
	uuidConnectedStr, err := middleware.ExtractTokenUUID(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	userToSend := ctx.Param("uuid")
	body := new(sendMoneyRequest)
	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	value := body.Value
	if value == 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": "value cannot be nul"})
		return
	}
	userConnected, err := h.GetUserByuuid(uuidConnectedStr)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	if err := middleware.IsTruePassword(userConnected.Password, body.Passwords); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	userRecepteur, err := h.GetUserByuuid(userToSend)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	// check si l'envoyeur essayent d'envoyer plus d'argent que ce qu'il en a
	if value > userConnected.Moneys {
		err := fmt.Errorf("it's impossible to send money of value %v , because money in your account %v", value, userConnected.Moneys)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	// check si l'userConnecter est la meme que celui qui il essaye d'envoyer de l'argent
	if uuidConnectedStr == userRecepteur.UUID {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": " you cannot send money to your self "})
		return
	}
	userConnected.Moneys = userConnected.Moneys - value
	userRecepteur.Moneys = (userRecepteur.Moneys + value)

	h.DB.Save(userRecepteur)
	h.DB.Save(userConnected)

	moneyTransaction, err := h.dbManipulationSendMoney(userConnected, userRecepteur, body)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"moneyTransaction": &moneyTransaction,
	})
}

func (h handler) dbManipulationSendMoney(userConnected, userRecepteur *models.User, body *sendMoneyRequest) (*models.Money, error) {
	var moneyTransaction models.Money
	n, found := slices.BinarySearch(userConnected.BlockedAcc, userRecepteur.AppUserName)
	if found {
		return nil, fmt.Errorf("user with uuid %s already blocked position : %v all blocekdAccount: %v", userRecepteur.UUID, n, userConnected.BlockedAcc)
	}
	resume := fmt.Sprintf("%v just send %v to %v now :%v ", userConnected.AppUserName, body.Value, userRecepteur.AppUserName, time.Now().Format(time.RFC1123Z))
	if res := h.DB.Where("send_by = ? AND sent_to = ?", userConnected.UUID, userRecepteur.UUID).First(&moneyTransaction); res.Error != nil {
		fmt.Printf("no transaction betwen %s and %s creating a new one.....", userConnected.UUID, userRecepteur.UUID)
		moneyTransaction.ID = uuid.New()
		moneyTransaction.Resume = resume
		moneyTransaction.Totals = body.Value
		moneyTransaction.SendBy = userConnected.UUID
		moneyTransaction.SentTo = userRecepteur.UUID
		moneyTransaction.SendByImg = userConnected.Image
		moneyTransaction.SendToImg = userRecepteur.Image
		moneyTransaction.SentToName = userRecepteur.AppUserName
		moneyTransaction.MoneyTransite = append(moneyTransaction.MoneyTransite, body.Value)
		result := h.DB.Create(moneyTransaction)
		if result.Error != nil {
			return nil, fmt.Errorf("creation raw %v", result.Error)
		}
		return &moneyTransaction, nil
	}
	fmt.Println("transaction already exist")
	moneyTransaction.Resume = resume
	moneyTransaction.MoneyTransite = append(moneyTransaction.MoneyTransite, body.Value)
	moneyTransaction.Totals = getTotals(moneyTransaction.MoneyTransite)
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
	res := h.DB.Where("uuid = ? OR app_user_name = ?", userReq, userReq).First(&users)
	if res.Error != nil {
		return nil, fmt.Errorf(" %v: uuid or AppUserName notFound", userReq)
	}
	return &users, nil
}
