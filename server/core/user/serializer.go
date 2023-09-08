package user

type UserSerializer interface {
	Decode(input []byte) (*User, error)
	Encode(input *User) ([]byte, error)
}
