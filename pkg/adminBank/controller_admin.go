package adminbank

// controller admin , code who all the admins action
import (
	"github.com/RazanakotoMandresy/bank-app-aout/backend/pkg/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func AdminRoutes(router *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}
	routes := router.Group("/api/v1/admin")
	routes.GET("/search", h.SearchBanks)
	routes.POST("/registerAnAdmin", h.CreateAdminAccount)
	routes.POST("/loginAsAdmin", h.LoginAdmin)
	routes.POST("/createBank", middleware.RequireAuth, h.CreateBank)
	routes.GET("/getBank", middleware.RequireAuth, h.GetBankLogAdmin)
	routes.GET("/getAdminInfo", middleware.RequireAuth, h.GetAdminInfo)
}
	