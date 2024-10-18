package chatrealtimes

import (
	"log"
	"net/http"
	"time"

	"github.com/RazanakotoMandresy/bank-app-aout/backend/pkg/common/models"
	"github.com/RazanakotoMandresy/bank-app-aout/backend/pkg/middleware"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h handler) handleWebSocket(ctx *gin.Context) {
	uuidToSend := ctx.Param("uuid")
	uuidSendBy, err := middleware.ExtractTokenUUID(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"err": err.Error()})
		return
	}
	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Println("Failed to set websocket upgrade: ", err)
		return
	}
	defer conn.Close()

	for {
		// Lire le message du client
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		newMessage := models.Chat{Content: string(message), SentTo: uuidToSend, SendBy: uuidSendBy, ID: uuid.New() , Created_at: time.Now().Local()}
		h.DB.Create(&newMessage)
		// Echo du message reçu au client
		err = conn.WriteMessage(messageType, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}
