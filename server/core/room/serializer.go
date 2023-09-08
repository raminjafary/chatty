package room

type RoomSerializer interface {
	Decode(input []byte) (*Room, error)
	Encode(input *Room) ([]byte, error)
	EncodeAll(input []*Room) ([]byte, error)
}
