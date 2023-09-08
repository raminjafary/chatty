package room

type RoomService interface {
	CreateRoom(req *Room) (*Room, error)
	GetRooms() ([]*Room, error)
}
