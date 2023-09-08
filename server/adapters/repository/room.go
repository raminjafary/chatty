package repository

import (
	"errors"
	"fmt"
	"server/core/room"

	"github.com/google/uuid"
)

func (rr *DB) CreateRoom(u *room.Room) (*room.Room, error) {
	newRoom := &room.Room{}
	req := rr.db.First(&newRoom, "roomName = ?", u.RoomName)

	if req.RowsAffected != 0 {
		return nil, errors.New("room already exists")
	}

	newRoom = &room.Room{
		ID:       uuid.New().String(),
		RoomName: u.RoomName,
		UserId:   u.UserId,
	}

	req = rr.db.Create(&newRoom)

	if req.RowsAffected == 0 {
		return nil, fmt.Errorf("room not saved: %v", req.Error)
	}

	return newRoom, nil
}

func (rr *DB) GetRooms() ([]*room.Room, error) {
	room := make([]*room.Room, 0)

	req := rr.db.First(&room)

	if req.RowsAffected == 0 {
		return nil, errors.New("user not found")
	}

	return room, nil
}
