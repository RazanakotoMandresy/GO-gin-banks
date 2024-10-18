package chatrealtimes

import (
	"net/http"

	"github.com/RazanakotoMandresy/bank-app-aout/backend/pkg/middleware"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"gorm.io/gorm"
)

type MessageRequest struct {
	Content string
	SendBy  string
	SentTo  string
}
type handler struct {
	DB *gorm.DB
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func ChatTransaction(router *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}
	routes := router.Group("/api/v1/chat")
	routes.GET("/", middleware.RequireAuth, h.GetUsrMsgSent)
	routes.GET("/ws/:uuid", middleware.RequireAuth, h.handleWebSocket)
	// ireo izay nalefa
	routes.GET("/:uuid", middleware.RequireAuth, h.GetAllMessageSent)
	routes.GET("/rec/:uuid", middleware.RequireAuth, h.ReceiveMessage)
}
