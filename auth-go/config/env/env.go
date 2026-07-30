package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func load() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading dotenv file")
	}

}

func GetString(key string, fallback string) string {
	load()
	value, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}
	return value
}

func GetInt(key string, fallback int) int {
	load()
	value, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}
	valueInt, err := strconv.Atoi(value)
	if err != nil {
		fmt.Println("Error converting env variable to int")
		return fallback
	}
	return valueInt
}


func GetBool(key string, fallback bool) bool {
	load()
	value, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}
	valueBool, err := strconv.ParseBool(value)
	if err != nil {
		fmt.Println("Error converting env variable to bool")
		return fallback
	}
	return valueBool
}
