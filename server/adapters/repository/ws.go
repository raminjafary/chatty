package repository

import (
	"github.com/gorilla/websocket"
)

type WS struct {
	upgrader *websocket.Upgrader
}

func NewWS(upgrader *websocket.Upgrader) *WS {
	return &WS{
		upgrader,
	}
}
