package utils

import (
	config "auth-go/config/env"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(email string, id string) (string, error) {
	secretKey := config.GetString("JWT_SECRET", "TOKEN")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"email": email, "id": id, "exp": time.Now().Add(time.Hour * 24).Unix()})
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		log.Fatal(err)
	}
	return tokenString, err

}
