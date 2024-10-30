package money

import (
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/RazanakotoMandresy/go-gin-banks/pkg/common/models"
	"github.com/RazanakotoMandresy/go-gin-banks/pkg/middleware"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h handler) SendMoney(ctx *gin.Context) {
	// extraction de l'uuid depuis le bearer
	body := new(sendMoneyRequest)
	uuidConnectedStr, err := middleware.ExtractTokenUUID(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	value := body.Value
	if value == 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": "value cannot be nul"})
		return
	}
	userConnected, err := middleware.GetUserUUID(h.DB, uuidConnectedStr)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	if err := middleware.IsTruePassword(userConnected.Password, body.Passwords); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	// ctx params userTo send money UUID
	userRecepteur, err := middleware.GetUserUUID(h.DB, ctx.Param("uuid"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	// check if user have less money in his account than he try to send
	if value > userConnected.Moneys {
		err := fmt.Errorf("it's impossible to send money of value %v , because money in your account %v", value, userConnected.Moneys)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	// if user try ti send money to himself
	if uuidConnectedStr == userRecepteur.UUID {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": " you cannot send money to your self "})
		return
	}

	userConnected.Moneys = userConnected.Moneys - value
	userRecepteur.Moneys = (userRecepteur.Moneys + value)

	h.DB.Save(userRecepteur)
	h.DB.Save(userConnected)

	moneyTransaction, err := h.dbManipulationSendMoney(body, userConnected, userRecepteur)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"moneyTransaction": &moneyTransaction,
	})
}

func (h handler) dbManipulationSendMoney(body *sendMoneyRequest, userConnected, userRecepteur *models.User) (*models.Money, error) {
	var moneyTransaction models.Money
	n, found := slices.BinarySearch(userConnected.BlockedAcc, userRecepteur.AppUserName)
	if found {
		return nil, fmt.Errorf("user with uuid %s already blocked position : %v all blocekdAccount: %v", userRecepteur.UUID, n, userConnected.BlockedAcc)
	}
	resume := fmt.Sprintf("%v just send %v to %v now :%v ", userConnected.AppUserName, body.Value, userRecepteur.AppUserName, time.Now().Format(time.RFC1123Z))
	fmt.Printf("no transaction betwen %s and %s creating a new one.....", userConnected.UUID, userRecepteur.UUID)
	moneyTransaction.ID = uuid.New()
	moneyTransaction.Resume = resume
	moneyTransaction.Totals = body.Value
	moneyTransaction.SendBy = userConnected.UUID
	moneyTransaction.SentTo = userRecepteur.UUID
	moneyTransaction.SendByImg = userConnected.Image
	moneyTransaction.SendToImg = userRecepteur.Image
	moneyTransaction.SentToName = userRecepteur.AppUserName
	result := h.DB.Create(moneyTransaction)
	if result.Error != nil {
		return nil, fmt.Errorf("creation raw %v", result.Error)
	}
	return &moneyTransaction, nil
}

// // user req que se soit uuid na appUserName
// func (h handler) GetUserByuuid(userReq string) (*models.User, error) {
// 	var users models.User
// 	res := h.DB.Where("uuid = ? OR app_user_name = ?", userReq, userReq).First(&users)
// 	if res.Error != nil {
// 		return nil, fmt.Errorf(" %v: uuid or AppUserName notFound", userReq)
// 	}
// 	return &users, nil
// }
