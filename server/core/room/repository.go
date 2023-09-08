package room

type RoomRepository interface {
	CreateRoom(room *Room) (*Room, error)
	GetRooms() ([]*Room, error)
}
