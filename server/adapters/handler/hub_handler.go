package handler

import (
	"log"
	"net/http"
	"server/core/hub"
	"server/core/member"
	"server/core/message"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type HubHandler interface {
	JoinRoom(*gin.Context)
}

type hubHandler struct {
	hubService hub.HubService
	hub        *hub.Hub
}

func NewHubHandler(hubService hub.HubService) HubHandler {
	return &hubHandler{hubService: hubService}
}

func (h *hubHandler) JoinRoom(c *gin.Context) {

	roomID := c.Param("roomId")
	userId := c.Query("userId")
	username := c.Query("username")

	m, msg, err := h.hubService.JoinRoom(c, roomID, userId, username)

	if err != nil {
		http.Error(c.Writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	h.hub.Register <- m
	h.hub.Broadcast <- msg

	go writeMessage(m)
	go readMessage(h.hub, m)
}

func writeMessage(c *member.Member) {
	defer func() {
		c.Conn.Close()
	}()

	for {
		message, ok := <-c.Message
		if !ok {
			return
		}

		c.Conn.WriteJSON(message)
	}
}

func readMessage(hub *hub.Hub, c *member.Member) {
	defer func() {
		hub.Unregister <- c
		c.Conn.Close()
	}()

	for {
		_, m, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		println(string(m))

		msg := &message.Message{
			Content:  string(m),
			RoomID:   c.RoomID,
			Username: c.Username,
		}

		hub.Broadcast <- msg
	}
}
