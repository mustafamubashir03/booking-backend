package services

import (
	db "auth-go/db/repositories"
	"auth-go/dto"
	"auth-go/models"
	"auth-go/utils"
	"fmt"
	"log"
)

type UserService interface {
	CreateUser(name string, email string, password string) error
	GetUserById(id string) (*models.User, error)
	GetAllUsers() ([]models.User, error)
	DeleteUserById(id string) error
	LoginUser(payload *dto.LoginUserRequestDTO) (string, error)
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

func (u *UserServiceImp) GetUserById(id string) (*models.User, error) {
	return u.userRepository.GetById(id)
}

func (u *UserServiceImp) GetAllUsers() ([]models.User, error) {
	return u.userRepository.GetAll()
}

func (u *UserServiceImp) DeleteUserById(id string) error {
	return u.userRepository.DeleteById(id)
}

func (u *UserServiceImp) LoginUser(payload *dto.LoginUserRequestDTO) (string, error) {
	userFound, err := u.userRepository.GetByEmail(payload.Email)
	if err != nil {
		return "", err
	}
	isMatch := utils.CompareHashedPassword([]byte(userFound.Password), payload.Password)
	if !isMatch {
		return "", fmt.Errorf("password does not match")
	}
	jwtToken, err := utils.GenerateJWT(userFound.Email, string(userFound.Id))
	if err != nil {
		return "", err
	}
	return jwtToken, nil

}
