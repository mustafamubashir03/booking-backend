package middlewares

import (
	"auth-go/utils"
	"context"
	"net/http"
)

func ValidateRequest(payload any) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if err := utils.ReadJsonBody(r, payload); err != nil {
				utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "JSON parse error.", err)
				return
			}

			err := utils.Validate.Struct(payload)

			if err != nil {
				utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Validation failed.", err)
				return
			}
			reqCtx := r.Context()
			ctx := context.WithValue(reqCtx, "payload", payload)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
