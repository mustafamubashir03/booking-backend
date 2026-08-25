package middlewares

import (
	"context"
	"net/http"

	"auth-go/utils"
)

type contextKey string

const RequestDTOKey contextKey = "requestDTO"

func ValidateRequest(payload any) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			if err := utils.ReadJsonBody(r, payload); err != nil {
				utils.WriteJsonErrorResponse(
					w,
					http.StatusBadRequest,
					"JSON parse error.",
					err,
				)
				return
			}

			if err := utils.Validate.Struct(payload); err != nil {
				utils.WriteJsonErrorResponse(
					w,
					http.StatusBadRequest,
					"Validation failed.",
					err,
				)
				return
			}

			ctx := context.WithValue(
				r.Context(),
				RequestDTOKey,
				payload,
			)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
