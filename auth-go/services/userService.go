package services

import (
	db "auth-go/db/repositories"
	"auth-go/utils"

	"log"
)

type UserService interface {
	CreateUser(name string, email string, password string) error
	GetUserById()
	GetAllUsers()
	DeleteUserById()
	LoginUser(email string, password string) string
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

func (u *UserServiceImp) CreateUser(name string, email string, password string) error {
	hashedPassword, err := utils.CreateHashPassword(password)
	if err != nil {
		log.Fatal("error in creating hash password")
	}
	u.userRepository.Create(name, email, string(hashedPassword))
	return nil
}

func (u *UserServiceImp) GetUserById() {
	u.userRepository.GetById()
}

func (u *UserServiceImp) GetAllUsers() {
	u.userRepository.GetAll()
}

func (u *UserServiceImp) DeleteUserById() {
	u.userRepository.DeleteById()
}

func (u *UserServiceImp) LoginUser(email string, password string) string {
	userFound, err := u.userRepository.GetByEmail(email)
	if err != nil {
		log.Fatal("error in getting user by email")
	}
	isMatch := utils.CompareHashedPassword([]byte(userFound.Password), password)
	if !isMatch {
		log.Fatal("password does not match")
	}
	jwtToken := utils.GenerateJWT(userFound.Username)
	return jwtToken

}
