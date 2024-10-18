package epargne

import (
	"github.com/RazanakotoMandresy/bank-app-aout/backend/pkg/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func EpargneTransaction(router *gin.Engine, db *gorm.DB) {

	h := &handler{
		DB: db,
	}
	routes := router.Group("/api/v1/epargne")
	routes.POST("/createEpargne", middleware.RequireAuth, h.CreateEpargne)
}
