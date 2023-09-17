package hub

import (
	"server/core/member"
	"server/core/message"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type hubService struct {
	hubRepo HubRepository
}

func NewHubService(hubRepo HubRepository) HubService {
	return &hubService{
		hubRepo,
	}
}

func (h *hubService) JoinRoom(c *gin.Context, roomID, userId, username string) (*member.Member, *message.Message, error) {

	conn, err := h.hubRepo.JoinRoom(c)

	if err != nil {
		return nil, nil, err
	}

	m := &member.Member{
		Conn:     conn,
		Message:  make(chan *message.Message, 10),
		ID:       uuid.New().String(),
		RoomID:   roomID,
		Username: username,
	}

	msg := &message.Message{
		Content:  "A new user has joined the room",
		RoomID:   roomID,
		Username: username,
	}

	return m, msg, nil

}

func (h *hubService) CreateHub() *Hub {
	return h.hubRepo.CreateHub()
}

func (h *Hub) Run() {
	for {
		select {
		case m := <-h.Register:
			if _, ok := h.Rooms[m.RoomID]; ok {
				r := h.Rooms[m.RoomID]

				if _, ok := r.Members[m.RoomID]; !ok {
					r.Members[m.RoomID] = m
				}
			}
		case m := <-h.Unregister:
			if _, ok := h.Rooms[m.RoomID]; ok {
				if _, ok := h.Rooms[m.RoomID].Members[m.ID]; ok {
					if len(h.Rooms[m.RoomID].Members) != 0 {
						h.Broadcast <- &message.Message{
							Content:  "user left the chat",
							RoomID:   m.RoomID,
							Username: m.Username,
						}
					}

					delete(h.Rooms[m.RoomID].Members, m.ID)
					close(m.Message)
				}
			}
		case msg := <-h.Broadcast:
			if _, ok := h.Rooms[msg.RoomID]; ok {
				for _, m := range h.Rooms[msg.RoomID].Members {
					m.Message <- msg
				}
			}
		}
	}
}
