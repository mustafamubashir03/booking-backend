package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func CompareHashedPassword(hashedPassword []byte, password string) bool {
	err := bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	return err == nil
}
