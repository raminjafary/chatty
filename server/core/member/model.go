package member

import (
	"server/core/message"
)

type Member struct {
	Message  chan *message.Message
	ID       string `json:"id"`
	RoomID   string `json:"roomId"`
	Username string `json:"username"`
}
