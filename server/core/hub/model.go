package hub

import (
	"server/core/member"
	"server/core/message"
	"server/core/room"
)

type Hub struct {
	Rooms      map[string]*room.Room
	Register   chan *member.Member
	Unregister chan *member.Member
	Broadcast  chan *message.Message
}
