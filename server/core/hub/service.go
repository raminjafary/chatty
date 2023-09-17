package hub

import (
	"server/core/member"
	"server/core/message"

	"github.com/gin-gonic/gin"
)

type HubService interface {
	CreateHub() *Hub
	JoinRoom(c *gin.Context, roomId, userId, username string) (*member.Member, *message.Message, error)
}
