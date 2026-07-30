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

func (u *UserController) RegisterUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("user controller hit")
	u.userService.CreateUser()
	w.Write([]byte("register successful"))
}
