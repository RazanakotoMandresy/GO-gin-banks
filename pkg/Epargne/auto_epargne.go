package epargne

import (
	"errors"
	"fmt"
	"time"

	"github.com/RazanakotoMandresy/go-gin-banks/pkg/common/models"
	"github.com/RazanakotoMandresy/go-gin-banks/pkg/middleware"
	"github.com/google/uuid"
)

func AutoEpargne(h Handler) error {
	var epargnes []models.Epargne
	fmt.Println("cronjob executed", time.Now().Format(time.DateTime))
	getEpargne := h.DB.Find(&epargnes)
	if getEpargne.Error != nil {
		return fmt.Errorf("err on get all epargnes: %v", getEpargne.Error)

	}
	for _, epargne := range epargnes {
		user, err := middleware.User.User(middleware.User{UuidToFind: epargne.OwnerUUID, Db: h.DB})
		if err != nil {
			return err
		}
		if time.Now().Day() == int(epargne.DayPerMounth) {
			if epargne.Type != "economies" {
				if err := economiesCase(h, user, &epargne); err != nil {
					return err
				}
				return err
			}
		}
	}
	return nil
}

func economiesCase(h Handler, user *models.User, epargne *models.Epargne) error {
	user.Moneys = (user.Moneys - epargne.Value)
	if res := h.DB.Save(&user); res.Error != nil {
		return errors.New("Error when updating user money" + res.Error.Error())
	}
	if createRes := h.DB.Create(&models.EpargneResume{
		ID:            uuid.New(),
		Type:          epargne.Type,
		ResumeMessage: fmt.Sprintf("epargne just got created : value %v , day %v , sent_to %s , owner %s", epargne.Value, epargne.DayPerMounth, epargne.Sent_to, epargne.OwnerUUID),
		Created_at:    time.Now(),
	}); createRes.Error != nil {
		return errors.New("Error occuring creating the epargne resume " + createRes.Error.Error())
	}
	return nil
}
