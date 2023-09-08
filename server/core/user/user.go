package user

import "errors"

var (
	ErrUsertNotFound = errors.New("User Not Found")
	ErrUserInvalid   = errors.New("User Invalid")
)

type userService struct {
	userRepo UserRepository
}

func NewUserService(userRepo UserRepository) UserService {
	return &userService{
		userRepo,
	}
}

func (r *userService) CreateUser(user *User) (*User, error) {
	return r.userRepo.CreateUser(user)
}

func (r *userService) GetUserById(id string) (*User, error) {
	return r.userRepo.GetUserById(id)
}
