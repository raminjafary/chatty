package repository

import (
	"errors"
	"fmt"
	"server/core/user"
	"time"

	"github.com/google/uuid"
)

func (ur *DB) CreateUser(u *user.User) (*user.User, error) {
	newUser := &user.User{}
	req := ur.db.First(&newUser, "email = ?", u.Email)

	if req.RowsAffected != 0 {
		return nil, errors.New("user already exists")
	}

	newUser = &user.User{
		ID:       uuid.New().String(),
		Email:    u.Email,
		Password: u.Password,
		Username: u.Username,
	}

	req = ur.db.Create(&newUser)

	if req.RowsAffected == 0 {
		return nil, fmt.Errorf("user not saved: %v", req.Error)
	}

	return newUser, nil
}

func (ur *DB) GetUserById(id string) (*user.User, error) {
	user := &user.User{}
	cachekey := user.ID
	err := ur.cache.Get(cachekey, &user)
	if err == nil {
		return user, nil
	}

	req := ur.db.First(&user, "id = ? ", id)

	if req.RowsAffected == 0 {
		return nil, errors.New("user not found")
	}

	err = ur.cache.Set(cachekey, user, time.Minute*10)

	if err != nil {
		fmt.Printf("Error storing user in cache: %v", err)
	}
	return user, nil
}
