package main

import (
	"fmt"
	"log"
	"os"

	epargne "github.com/RazanakotoMandresy/go-gin-banks/pkg/Epargne"
	adminbank "github.com/RazanakotoMandresy/go-gin-banks/pkg/adminBank"
	chatrealtimes "github.com/RazanakotoMandresy/go-gin-banks/pkg/chatRealtimes"
	"github.com/RazanakotoMandresy/go-gin-banks/pkg/common/db"
	"github.com/RazanakotoMandresy/go-gin-banks/pkg/money"
	"github.com/RazanakotoMandresy/go-gin-banks/pkg/user"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"
)

func main() {
	// cron for epargne
	godotenv.Load(".env")
	port := os.Getenv("PORT")
	dbUrl := os.Getenv("DB_URL")
	dbHandler := db.Init(dbUrl)
	newCron := cron.New()
	newCron.AddFunc("@daily", func() {
		if err := epargne.AutoEpargne(epargne.Handler{DB: dbHandler}); err != nil {
			fmt.Println("error on auto epargne ... :", err)
		}
	})
	newCron.Start()
	defer newCron.Stop()
	router := gin.Default()
	gin.SetMode(gin.DebugMode)
	router.Use(CORSMiddleware())
	user.RegisterRoutes(router, dbHandler)
	money.TransactionRoutes(router, dbHandler)
	adminbank.AdminRoutes(router, dbHandler)
	epargne.EpargneTransaction(router, dbHandler)
	// websocket ny chatrealtimes
	chatrealtimes.ChatTransaction(router, dbHandler)
	// dir misy amzao
	rootDir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	// serve depuis ou prendre les images et donne l'url
	router.Static("./upload", rootDir+"/upload")
	router.Static("./imgDef", rootDir+"/imgDef")
	if err := router.Run(port); err != nil {
		log.Fatal("an error occured during running the router", err)
	}

}
