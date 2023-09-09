package message

type messageService struct {
	messageRepo MessageRepository
}

func NewMessageService(messageRepo MessageRepository) MessageService {
	return &messageService{
		messageRepo,
	}
}

func (m *messageService) CreateMessage(message *Message) (*Message, error) {
	return m.messageRepo.CreateMessage(message)
}
