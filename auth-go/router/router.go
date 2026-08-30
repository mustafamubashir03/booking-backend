package router

import (
	"auth-go/controllers"
	"auth-go/middlewares"
	"auth-go/utils"

	"github.com/go-chi/chi/v5"
)

type Router interface {
	Register(r chi.Router)
}

func SetupRouter(userRouter Router, roleRouter Router) *chi.Mux {
	chiRouter := chi.NewRouter()
	chiRouter.Use(middlewares.RequestLogger)
	chiRouter.Use(middlewares.RateLimiter)
	chiRouter.Handle("/fake-store/*", utils.ProxyToService("https://fakestoreapi.com/products", "/fake-store"))
	chiRouter.Get("/ping", controllers.PingHandler)
	userRouter.Register(chiRouter)
	roleRouter.Register(chiRouter)
	return chiRouter
}
