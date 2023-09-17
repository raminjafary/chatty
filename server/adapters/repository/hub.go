package repository

import (
	"server/core/hub"
	"server/core/member"
	"server/core/message"
	"server/core/room"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func (ws *WS) CreateHub() *hub.Hub {
	return &hub.Hub{
		Rooms:      make(map[string]*room.Room),
		Register:   make(chan *member.Member),
		Unregister: make(chan *member.Member),
		Broadcast:  make(chan *message.Message, 5),
	}
}

func (ws *WS) JoinRoom(c *gin.Context) (*websocket.Conn, error) {

	conn, err := ws.upgrader.Upgrade(c.Writer, c.Request, nil)

	if err != nil {
		return nil, err
	}

	return conn, nil

}
