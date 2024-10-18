package db

import (
	"log"

	"github.com/RazanakotoMandresy/bank-app-aout/backend/pkg/common/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// init the db

func Init(url string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Money{})
	db.AutoMigrate(&models.Admin{})
	db.AutoMigrate(&models.Bank{})
	db.AutoMigrate(&models.Epargne{})
	db.AutoMigrate(&models.Chat{})
	return db
}
