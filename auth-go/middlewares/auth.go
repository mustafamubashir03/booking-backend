package middlewares

import (
	config "auth-go/config/env"
	"auth-go/utils"
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth_header := r.Header.Get("Authorization")
		if auth_header == "" {
			http.Error(w, "Un authorized", http.StatusUnauthorized)
			return
		}
		if !strings.HasPrefix(auth_header, "Bearer ") {
			http.Error(w, "Authorization header must start with Bearer", http.StatusBadRequest)
			return
		}
		tokenString := strings.TrimPrefix(auth_header, "Bearer ")
		claims := jwt.MapClaims{}
		_, error := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.GetString("JWT_SECRET", "TOKEN")), nil
		})
		if error != nil {
			http.Error(w, "Invalid Token", http.StatusUnauthorized)
			return
		}
		userId, ok := claims["id"].(string)
		email, okEmail := claims["email"].(string)
		if !ok || !okEmail {
			utils.WriteJsonErrorResponse(w, http.StatusUnauthorized, "Invalid Token", error)
			return
		}
		fmt.Println("Authenticated with :" + email + " and id:" + userId)
		ctx := context.WithValue(r.Context(), "userId", userId)
		ctx = context.WithValue(ctx, "email", email)
		next.ServeHTTP(w, r.WithContext(ctx))

	})

}
