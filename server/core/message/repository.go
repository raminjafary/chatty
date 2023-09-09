package message

type MessageRepository interface {
	CreateMessage(*Message) (*Message, error)
}
