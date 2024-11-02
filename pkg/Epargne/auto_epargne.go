package epargne

import (
	"fmt"
	"time"

	"github.com/RazanakotoMandresy/go-gin-banks/pkg/common/models"
	"github.com/RazanakotoMandresy/go-gin-banks/pkg/middleware"
	"github.com/google/uuid"
)

func (h Handler) AutoEpargne() {
	var epargnes []models.Epargne
	getEpargne := h.DB.Find(&epargnes)
	if getEpargne.Error != nil {
		fmt.Printf("err on get all epargnes: %v", getEpargne.Error)
		return
	}
	for _, epargne := range epargnes {
		user, err := middleware.User.User(middleware.User{UuidToFind: epargne.OwnerUUID, Db: h.DB})
		if err != nil {
			return
		}
		if time.Now().Day() == int(epargne.DayPerMounth) {
			user.Moneys = (user.Moneys - epargne.Value)
			if res := h.DB.Save(&user); res.Error != nil {
				return
			}
			if createRes := h.DB.Create(&models.EpargneResume{
				ID:            uuid.New(),
				Type:          epargne.Type,
				ResumeMessage: fmt.Sprintf("epargne just got created : value %v , day %v , sent_to %s , owner %s", epargne.Value, epargne.DayPerMounth, epargne.Sent_to, epargne.OwnerUUID),
				Created_at:    time.Now(),
			}); createRes.Error != nil {
				fmt.Printf("Error occuring creating the epargne resume %v", err)
				return
			}
		}
	}
}
