package member

import (
	"server/core/message"

	"github.com/gorilla/websocket"
)

type Member struct {
	Conn     *websocket.Conn       `json:"omitempty"`
	Message  chan *message.Message `json:"omitempty"`
	ID       string                `json:"id"`
	RoomID   string                `json:"roomId"`
	Username string                `json:"username"`
}
