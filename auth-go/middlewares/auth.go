package middlewares

import (
	config "auth-go/config/env"
	dbConfig "auth-go/db"
	dbRepository "auth-go/db/repositories"
	"auth-go/utils"
	"context"
	"fmt"
	"net/http"
	"strconv"
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
		_, error := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
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
		// fmt.Printf("CLAIMS ID: %v | BYTES: %v\n", claims["id"], []byte(userId))
		// fmt.Printf("USER ID: %v | BYTES: %v\n", userId, []byte(userId))
		// fmt.Printf("EMAIL: %v\n", email)
		fmt.Println("Authenticated with :" + email + " and id:" + userId)
		ctx := context.WithValue(r.Context(), "userId", userId)
		ctx = context.WithValue(ctx, "email", email)
		next.ServeHTTP(w, r.WithContext(ctx))

	})

}

func RequireAllRolesMiddleware(roles ...string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userId := r.Context().Value("userId").(string)
			dbConn, dbErr := dbConfig.SetupDB()
			if dbErr != nil {
				utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Database error", dbErr)
				return
			}
			userRoleRepository := dbRepository.NewUserRoleRepository(dbConn)
			userIdInt, err := strconv.Atoi(userId)
			if err != nil {
				utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Invalid User ID", err)
				return
			}
			exists, err := userRoleRepository.HasAllRoles(userIdInt, roles)
			if err != nil {
				utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Error checking roles", err)
				return
			}
			if !exists {
				utils.WriteJsonErrorResponse(w, http.StatusUnauthorized, "Unauthorized", nil)
				return
			}
			next.ServeHTTP(w, r)
		})
	}

}
func RequireAnyRoleMiddleware(roles ...string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userId := r.Context().Value("userId").(string)
			dbConn, dbErr := dbConfig.SetupDB()
			if dbErr != nil {
				utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Database error", dbErr)
				return
			}
			userRoleRepository := dbRepository.NewUserRoleRepository(dbConn)
			userIdInt, err := strconv.Atoi(userId)
			fmt.Printf("USER ID INT: %v | %T\n", userIdInt, userIdInt)
			if err != nil {
				utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Invalid User ID", err)
				return
			}
			exists, err := userRoleRepository.HasAnyRole(userIdInt, roles)
			if err != nil {
				utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Error checking roles", err)
				return
			}
			if !exists {
				utils.WriteJsonErrorResponse(w, http.StatusUnauthorized, "Unauthorized", nil)
				return
			}
			next.ServeHTTP(w, r)
		})
	}

}
