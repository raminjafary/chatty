package hub

type RoomSerializer interface {
	Decode(input []byte) (*Hub, error)
	Encode(input *Hub) ([]byte, error)
}
