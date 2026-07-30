package services

import (
	db "auth-go/db/repositories"
	"fmt"
)

type UserService interface {
	CreateUser()
}

type UserServiceImp struct {
	userRepository db.UserRepository
}

func NewUserService(_userRepository db.UserRepository) UserService {
	userService := &UserServiceImp{
		userRepository: _userRepository,
	}
	return userService
}

func (u *UserServiceImp) CreateUser() {
	fmt.Println("user service hit")
	u.userRepository.Create()

}
