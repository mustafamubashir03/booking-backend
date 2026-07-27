package main

import (
	"auth-go/app"
	config "auth-go/config/env"
	"fmt"
)

func main() {
	fmt.Println("hello world")
	port := config.GetString("PORT", ":8080")
	cfg := app.NewConfig(port)
	app := app.NewApp(cfg)

	app.Run()
}
