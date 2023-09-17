package json

import (
	"encoding/json"

	"server/core/room"
	"server/core/user"

	"github.com/pkg/errors"
)

type User struct{}
type Room struct {
}

func (r *User) Decode(input []byte) (*user.User, error) {
	user := &user.User{}
	if err := json.Unmarshal(input, user); err != nil {
		return nil, errors.Wrap(err, "serializer.Uesr.Decode")
	}
	return user, nil
}

func (r *User) Encode(input *user.User) ([]byte, error) {
	rawMsg, err := json.Marshal(input)
	if err != nil {
		return nil, errors.Wrap(err, "serializer.User.Encode")
	}
	return rawMsg, nil
}

func (r *Room) Decode(input []byte) (*room.Room, error) {
	room := &room.Room{}
	if err := json.Unmarshal(input, room); err != nil {
		return nil, errors.Wrap(err, "serializer.Room.Decode")
	}
	return room, nil
}

func (r *Room) Encode(input *room.Room) ([]byte, error) {
	rawMsg, err := json.Marshal(input)
	if err != nil {
		return nil, errors.Wrap(err, "serializer.Room.Encode")
	}
	return rawMsg, nil
}

func (r *Room) EncodeAll(input []*room.Room) ([]byte, error) {
	rawMsg, err := json.Marshal(input)
	if err != nil {
		return nil, errors.Wrap(err, "serializer.Room.Encode")
	}
	return rawMsg, nil
}
