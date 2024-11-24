package epargne

import (
	"github.com/RazanakotoMandresy/go-gin-banks/pkg/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func EpargneTransaction(router *gin.Engine, db *gorm.DB) {
	h := &Handler{
		DB: db,
	}
	routes := router.Group("/api/v1/epargne")
	routes.POST("/createEpargne", middleware.RequireAuth, h.CreateEpargne)
	// getMyEpargne is only for non AutoSend and economies
	routes.GET("/", middleware.RequireAuth, h.GetAllMyEpargnes)
	routes.GET("/:epargneUUID", middleware.RequireAuth, h.GetMoneyEpargne)
	routes.DELETE("/:epargneUUID", middleware.RequireAuth, h.DeleteEpargne)
}
