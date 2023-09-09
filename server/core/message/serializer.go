package message

type MessageSerializer interface {
	Decode(input []byte) (*Message, error)
	Encode(input *Message) ([]byte, error)
}
