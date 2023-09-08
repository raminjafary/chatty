package room

type roomService struct {
	roomRepo RoomRepository
}

func NewRoomService(roomRepo RoomRepository) RoomService {
	return &roomService{
		roomRepo,
	}
}

func (r *roomService) CreateRoom(room *Room) (*Room, error) {
	return r.roomRepo.CreateRoom(room)
}

func (r *roomService) GetRooms() ([]*Room, error) {
	return r.roomRepo.GetRooms()
}
