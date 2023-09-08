package user

type UserService interface {
	CreateUser(user *User) (*User, error)
	GetUserById(id string) (*User, error)
}
