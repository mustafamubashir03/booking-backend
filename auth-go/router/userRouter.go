package router

import (
	"auth-go/controllers"

	"github.com/go-chi/chi/v5"
)

type UserRouter struct {
	userController *controllers.UserController
}

func NewUserRouter(_userController *controllers.UserController) Router {
	userRouter := &UserRouter{
		userController: _userController,
	}
	return userRouter
}

func (userRouter *UserRouter) Register(r chi.Router) {
	r.Post("/register", userRouter.userController.RegisterUser)

}
