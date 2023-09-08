package handler

import (
	"io"
	"net/http"
	"server/core/room"
	jsonS "server/serializer/json"

	"github.com/gin-gonic/gin"
)

type RoomHandler interface {
	CreateRoom(*gin.Context)
	GetRooms(*gin.Context)
}

type roomHandler struct {
	roomService room.RoomService
}

func NewRoomHandler(roomService room.RoomService) RoomHandler {
	return &roomHandler{roomService: roomService}
}

func (h *roomHandler) serializer(contentType string) room.RoomSerializer {
	return &jsonS.Room{}
}

func (h *roomHandler) CreateRoom(c *gin.Context) {
	contentType := c.Request.Header.Get("Content-Type")
	requestBody, err := io.ReadAll(c.Request.Body)

	if err != nil {
		http.Error(c.Writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	room, err := h.serializer(contentType).Decode(requestBody)

	if err != nil {
		http.Error(c.Writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	_, err = h.roomService.CreateRoom(room)

	if err != nil {
		http.Error(c.Writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	responseBody, err := h.serializer(contentType).Encode(room)

	if err != nil {
		http.Error(c.Writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	setupResponse(c, contentType, responseBody, http.StatusCreated)

}

func (h *roomHandler) GetRooms(c *gin.Context) {
	contentType := c.Request.Header.Get("Content-Type")

	room, err := h.roomService.GetRooms()

	if err != nil {
		http.Error(c.Writer, http.StatusText(http.StatusInternalServerError), http.StatusBadRequest)
		return
	}

	responseBody, err := h.serializer(contentType).EncodeAll(room)

	setupResponse(c, contentType, responseBody, http.StatusCreated)
}
