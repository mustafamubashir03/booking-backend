package app

import (
	"auth-go/router"
	"fmt"
	"net/http"
	"time"
)

type Config struct {
	Addr string
}
type Application struct {
	Config Config
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
	}
	return app
}

func (app *Application) Run() error {
	server := &http.Server{
		Addr:         app.Config.Addr,
		Handler:      router.SetupRouter(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Println("starting server on: ", app.Config.Addr)

	return server.ListenAndServe()

}
