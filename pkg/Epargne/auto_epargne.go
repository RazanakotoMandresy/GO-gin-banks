package epargne

import (
	"fmt"
	"log"

	"github.com/RazanakotoMandresy/go-gin-banks/pkg/common/models"
	"github.com/RazanakotoMandresy/go-gin-banks/pkg/middleware"
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
		user.Moneys = (user.Moneys - epargne.Value)
		if res := h.DB.Save(&user); res.Error != nil {
			return
		}

	}
	fmt.Println("soustract user money for the epargne ")

}
