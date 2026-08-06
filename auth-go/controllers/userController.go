package controllers

import (
	"auth-go/services"
	"fmt"
	"net/http"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(_userService services.UserService) *UserController {
	userController := &UserController{
		userService: _userService,
	}
	return userController

}

func (u *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("user controller hit")
	u.userService.CreateUser("shubham", "[EMAIL_ADDRESS]", "123456")
	w.Write([]byte("created user successful"))
}

func (u *UserController) GetUserById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("user controller hit")
	u.userService.GetUserById()
	w.Write([]byte("got user successful"))
}

func (u *UserController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("user controller hit")
	u.userService.GetAllUsers()
	w.Write([]byte("got all users successful"))
}

func (u *UserController) DeleteById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("user controller hit")
	u.userService.DeleteUserById()
	w.Write([]byte("deleted user successful"))
}
func (u *UserController) LoginUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("user controller hit")
	email := "[EMAIL_ADDRESS]"
	password := "123456"
	jwtToken := u.userService.LoginUser(email, password)
	w.Write([]byte(jwtToken))
}
