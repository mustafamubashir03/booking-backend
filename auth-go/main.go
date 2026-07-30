package main

import (
	"auth-go/app"
	config "auth-go/config/env"
	"auth-go/db"
	"fmt"
)

func main() {
	fmt.Println("hello world")
	port := config.GetString("PORT", ":8080")
	cfg := app.NewConfig(port)
	app := app.NewApp(cfg)
	db.SetupDB()
	app.Run()
}
