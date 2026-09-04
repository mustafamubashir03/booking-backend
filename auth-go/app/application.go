package app

import (
	"auth-go/controllers"
	dbConfig "auth-go/db"
	dbRepository "auth-go/db/repositories"
	"auth-go/router"
	"auth-go/services"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Config struct {
	Addr string
}
type Application struct {
	Config Config
	Store  dbRepository.Storage
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
		Store:  *dbRepository.NewStorage(),
	}
	return app
}

func (app *Application) Run() error {
	db, error := dbConfig.SetupDB()
	if error != nil {
		log.Fatalf("Error setting up database: %v", error)
	}
	userDb := dbRepository.NewUserRepository(db)
	userService := services.NewUserService(userDb)
	userController := controllers.NewUserController(userService)
	userRouter := router.NewUserRouter(userController)

	roleDb := dbRepository.NewRoleRepository(db)
	permissionDb := dbRepository.NewPermissionRepository(db)
	userRoleDb := dbRepository.NewUserRoleRepository(db)
	roleService := services.NewRoleService(roleDb, permissionDb, userRoleDb)
	roleController := controllers.NewRoleController(roleService)
	roleRouter := router.NewRoleRouter(roleController)

	server := &http.Server{
		Addr:         app.Config.Addr,
		Handler:      router.SetupRouter(userRouter, roleRouter),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Println("starting server on: ", app.Config.Addr)

	return server.ListenAndServe()

}
