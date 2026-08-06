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
	r.Get("/profile", userRouter.userController.GetUserById)
	r.Post("/sign-in", userRouter.userController.CreateUser)
	r.Get("/", userRouter.userController.GetAllUsers)
	r.Delete("/", userRouter.userController.DeleteById)
	r.Post("/login", userRouter.userController.LoginUser)
}
