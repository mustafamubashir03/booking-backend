package utils

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(username string) string {
	secretKey := []byte("secret-key")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"username": username, "exp": time.Now().Add(time.Hour * 24).Unix()})
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		log.Fatal(err)
	}
	return tokenString

}
