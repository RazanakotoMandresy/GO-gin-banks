package user

import (
	"errors"
	"fmt"
	"slices"

	"github.com/RazanakotoMandresy/go-gin-banks/pkg/common/models"
	"github.com/RazanakotoMandresy/go-gin-banks/pkg/middleware"
	"github.com/gin-gonic/gin"

	"net/http"
)

func (h handler) SettingUser(ctx *gin.Context) {
	body := new(SettingReq)
	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	// uuid uuid anze connecter
	uuid, err := middleware.ExtractTokenUUID(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"err": err.Error()})
		return
	}
	user, err := middleware.User.User(middleware.User{UuidToFind: uuid, Db: h.DB})
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"err": err.Error()})
		return
	}
	if body.RemoveAllEp {
		err := removeAllEpFunc(h, user.UUID)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
			return
		}
	}
	if body.DeleteMyAccount {
		err := deleMyAccount(h, user)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
			return
		}
	}
	if body.BlockAccount != "" {
		err := blockAccount(h, uuid, body.BlockAccount)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
			return
		}
	}
	if body.UnBlockAccount != "" {
		err := unBlockAccount(h, *user, body.UnBlockAccount)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		}
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"res": fmt.Sprintf("option %v clicked", body)})
}
func removeAllEpFunc(h handler, uuidUser string) error {
	var epargne models.Epargne
	res := h.DB.Delete(&epargne, "user_uuid = ?", uuidUser)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func deleMyAccount(h handler, user *models.User) error {
	res := h.DB.Delete(user, "uuid = ?", user.UUID)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func blockAccount(h handler, uuid, userBlockUUID string) error {
	// atao appUserName fa tsy uuid ny blockage
	userToBlock, err := middleware.User.User(middleware.User{UuidToFind: userBlockUUID, Db: h.DB})
	if err != nil {
		return err
	}
	user, err := middleware.User.User(middleware.User{UuidToFind: uuid, Db: h.DB})
	if err != nil {
		return err
	}
	// la logique ne fonctionne pas encore
	_, found := slices.BinarySearch(user.BlockedAcc, userToBlock.AppUserName)
	// une deuxieme appuye sur le block va causer un unblock
	if found {
		// appeles de unblock
		unBlockAccount(h, *user, userBlockUUID)
		return nil
	}
	if user.AppUserName == userToBlock.AppUserName {
		return errors.New("vous ne pouvez pas vous autobloquez")
	}
	user.BlockedAcc = append(user.BlockedAcc, userToBlock.AppUserName)
	res := h.DB.Save(&user)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

// pour un code plus safe de uuid string no naverina nalefa fa tsy tong de le user minyts
func unBlockAccount(h handler, user models.User, userUnblock string) error {
	userToUnblock, err := middleware.User.User(middleware.User{UuidToFind: userUnblock, Db: h.DB})
	// l'iuser ou l'on veut bloquer
	if err != nil {
		return err
	}
	blockedUser := []string{}
	for _, appUserName := range user.BlockedAcc {
		if appUserName == userToUnblock.AppUserName {
			continue
		}
		blockedUser = append(blockedUser, appUserName)
	}
	user.BlockedAcc = blockedUser
	h.DB.Save(&user)
	return nil
}
