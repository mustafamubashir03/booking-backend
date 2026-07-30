package app

import (
	"auth-go/controllers"
	db "auth-go/db/repositories"
	"auth-go/router"
	"auth-go/services"
	"fmt"
	"net/http"
	"time"
)

type Config struct {
	Addr string
}
type Application struct {
	Config Config
	Store  db.Storage
}

func NewConfig(addr string) Config {
	config := Config{
		Addr: addr,
	}
	return config
}

func NewApp(config Config) Application {
	app := Application{
		Config: config,
		Store:  *db.NewStorage(),
	}
	return app
}

func (app *Application) Run() error {
	userDb := db.NewUserRepository()
	userService := services.NewUserService(userDb)
	userController := controllers.NewUserController(userService)
	userRouter := router.NewUserRouter(userController)

	server := &http.Server{
		Addr:         app.Config.Addr,
		Handler:      router.SetupRouter(userRouter),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Println("starting server on: ", app.Config.Addr)

	return server.ListenAndServe()

}
