package money

import (
	"github.com/RazanakotoMandresy/go-gin-banks/pkg/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// models money
type handler struct {
	DB *gorm.DB
}

func TransactionRoutes(router *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}
	routes := router.Group("/api/v1/transaction")
	routes.GET("/historics", middleware.RequireAuth, h.HistoricTransaction)
	routes.PUT("/depot", middleware.RequireAuth, h.Depot)
	routes.PUT("/retrait", middleware.RequireAuth, h.Retrait)
	routes.POST("/:uuid", middleware.RequireAuth, h.SendMoney)
}
