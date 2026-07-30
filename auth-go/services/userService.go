package services

import (
	db "auth-go/db/repositories"
	"fmt"
)

type UserService interface {
	CreateUser()
	GetUserById()
	GetAllUsers()
	DeleteUserById()
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

func (u *UserServiceImp) GetUserById() {
	fmt.Println("getting user by id hit")
	u.userRepository.GetById()
}

func (u *UserServiceImp) GetAllUsers() {
	fmt.Println("getting all users hit")
	u.userRepository.GetAll()
}

func (u *UserServiceImp) DeleteUserById() {
	fmt.Println("deleting user by id hit")
	u.userRepository.DeleteById()
}
