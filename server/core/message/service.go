package message

type MessageService interface {
	CreateMessage(*Message) (*Message, error)
}
