package epargne

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/RazanakotoMandresy/go-gin-banks/pkg/common/models"
	"github.com/RazanakotoMandresy/go-gin-banks/pkg/middleware"
	"github.com/google/uuid"
)

func AutoEpargne(h Handler) error {
	var epargnes []models.Epargne
	// var userSentTo models.User
	fmt.Println("cronjob executed", time.Now().Format(time.DateTime))
	// get all epargne from the db
	if getEpargne := h.DB.Find(&epargnes); getEpargne.Error != nil {
		return fmt.Errorf("err on get all epargnes: %v", getEpargne.Error)
	}
	// get all the epargnes and sort them
	for _, epargne := range epargnes {
		user, err := middleware.User.User(middleware.User{UuidToFind: epargne.OwnerUUID, Db: h.DB})
		if err != nil {
			return err
		}
		// check day today and if the epargne is autosend mbola todo ny index ny logic false autosend
		if time.Now().Day() == int(epargne.DayPerMounth) && epargne.AutoSend && epargne.Type != "economies" && epargne.Type != "economie" {
			// handle logic economies ,dayPer month just day for the money to be soustract in the current user
			// the money will be send to himself but only if he click on the user on get my epargne and get the epargne with the specific uuid
			if err := autoEpargneCaseNoEconomie(user, &epargne, &h); err != nil {
				log.Default().Printf("error on epargne ")
			}
		}
		// case non auto send
		// non auto Send literaly another routes get Epargne
	}
	return nil
}
func autoEpargneCaseNoEconomie(user *models.User, epargne *models.Epargne, h *Handler) error {
	user.Moneys = (user.Moneys - epargne.Value)
	if res := h.DB.Save(&user); res.Error != nil {
		return errors.New("Error when updating user money" + res.Error.Error())
	}
	if createRes := h.DB.Create(&models.EpargneResume{
		ID:            uuid.New(),
		Type:          epargne.Type,
		ResumeMessage: fmt.Sprintf("value %v , day %v , sent_to %s , owner %s", epargne.Value, epargne.DayPerMounth, epargne.Sent_to, epargne.OwnerUUID),
		Created_at:    time.Now(),
		Value:         uint(epargne.Value),
	}); createRes.Error != nil {
		return errors.New("Error occuring creating the epargne resume " + createRes.Error.Error())
	}
	userToSend, err := middleware.User.User(middleware.User{UuidToFind: epargne.Sent_to, Db: h.DB})
	if err != nil {
		return err
	}
	userToSend.Moneys = user.Moneys + epargne.Value
	if errSaveSend := h.DB.Save(&userToSend); errSaveSend.Error != nil {
		return err
	}
	return nil
}
