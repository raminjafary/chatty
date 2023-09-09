package room

import "server/core/member"

type Room struct {
	ID       string                    `json:"id"`
	RoomName string                    `json:"roomName"`
	UserId   string                    `json:"userId"`
	Members  map[string]*member.Member `json:"members"`
}
