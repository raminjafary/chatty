package room

type RoomService interface {
	CreateRoom(room *Room) (*Room, error)
	GetRooms() ([]*Room, error)
}
