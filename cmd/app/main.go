package main

import (
	"fmt"
	"log"
	"os"

	epargne "github.com/RazanakotoMandresy/bank-app-aout/backend/pkg/Epargne"
	adminbank "github.com/RazanakotoMandresy/bank-app-aout/backend/pkg/adminBank"
	chatrealtimes "github.com/RazanakotoMandresy/bank-app-aout/backend/pkg/chatRealtimes"
	"github.com/RazanakotoMandresy/bank-app-aout/backend/pkg/common/db"
	"github.com/RazanakotoMandresy/bank-app-aout/backend/pkg/money"
	"github.com/RazanakotoMandresy/bank-app-aout/backend/pkg/user"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	router := gin.Default()
	gin.SetMode(gin.DebugMode)
	router.Use(CORSMiddleware())
	//
	godotenv.Load("./pkg/common/envs/.env")
	port := os.Getenv("PORT")
	dbUrl := os.Getenv("DB_URL")
	dbHandler := db.Init(dbUrl)
	//
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
