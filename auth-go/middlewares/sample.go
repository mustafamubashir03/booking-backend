package middlewares

import (
	"fmt"
	"net/http"
)

func RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Request Method: ", r.Method)
		fmt.Println("Request URL: ", r.URL)
		fmt.Println("Request Headers: ", r.Header)
		// fmt.Println("Request Body: ", r.Body)
		next.ServeHTTP(w, r)
	})
}
