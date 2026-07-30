package db

import (
	config "auth-go/config/env"
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

func SetupDB() (*sql.DB, error) {
	cfg := mysql.NewConfig()
	cfg.User = config.GetString("DATABASE_USER", "root")
	cfg.Passwd = config.GetString("DATABASE_PASSWORD", "root123")
	cfg.Addr = config.GetString("DATABASE_ADDRESS", "127.0.0.1:3306")
	cfg.DBName = config.GetString("DATABASE_NAME", "auth_db_dev")
	cfg.Net = "tcp"
	fmt.Println("connecting to db", cfg.DBName)
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		fmt.Println("error")
		log.Fatalf("Error opening database: %v", err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("error")
		log.Fatalf("Error pinging database: %v", err)
	}
	fmt.Println("Database connected successfully")
	return db, nil
}
