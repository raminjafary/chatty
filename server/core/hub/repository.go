package hub

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type HubRepository interface {
	CreateHub() *Hub
	JoinRoom(c *gin.Context) (*websocket.Conn, error)
}
