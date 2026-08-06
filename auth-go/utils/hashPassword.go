package utils

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func CreateHashPassword(password string) ([]byte, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return hashedPassword, nil
}
