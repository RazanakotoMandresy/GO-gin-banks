package user

import (
	"github.com/RazanakotoMandresy/bank-app-aout/backend/pkg/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}
	routes := router.Group("api/v1/user")

	routes.GET("/", middleware.RequireAuth, h.GetUsers)
	routes.GET("/logedUser", middleware.RequireAuth, h.getConnectedUser)
	routes.GET("/search", middleware.RequireAuth, h.SearchUser)
	// get single user apres recheche ou tuc du genres
	routes.GET("/:user", middleware.RequireAuth, h.getUser)
	routes.POST("/register", h.CreateUser)
	routes.POST("/login", h.Login)
	// change pp only (may i can fusion this with setting or smthg)
	routes.POST("/pp", middleware.RequireAuth, h.UserPP)
	routes.PATCH("/", middleware.RequireAuth, h.UpdateInfo)
	routes.PATCH("/setting", middleware.RequireAuth, h.SettingUser)
}
	