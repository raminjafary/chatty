package user

type UserRepository interface {
	CreateUser(user *User) (*User, error)
	GetUserById(id string) (*User, error)
}
