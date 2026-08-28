package router

import (
	"auth-go/controllers"
	"auth-go/dto"
	"auth-go/middlewares"

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
	r.With(middlewares.ValidateRequest(&dto.RegisterUserRequestDTO{})).Post("/sign-in", userRouter.userController.CreateUser)
	r.With(middlewares.ValidateRequest(&dto.LoginUserRequestDTO{})).Post("/login", userRouter.userController.LoginUser)
	r.With(middlewares.AuthMiddleware).Get("/", userRouter.userController.GetAllUsers)
	r.With(middlewares.AuthMiddleware).Get("/profile", userRouter.userController.GetUserById)
	r.With(middlewares.AuthMiddleware).Delete("/", userRouter.userController.DeleteById)
}
