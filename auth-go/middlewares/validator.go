package middlewares

import (
	"auth-go/utils"
	"context"
	"net/http"
	"reflect"
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

func ValidateParams(payload any) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			value := reflect.ValueOf(payload)

			if value.Kind() != reflect.Pointer || value.Elem().Kind() != reflect.Struct {
				utils.WriteJsonErrorResponse(
					w,
					http.StatusInternalServerError,
					"Invalid validation payload.",
					nil,
				)
				return
			}

			value = value.Elem()
			typeOfPayload := value.Type()

			for i := 0; i < value.NumField(); i++ {
				field := value.Field(i)
				structField := typeOfPayload.Field(i)

				paramName := structField.Tag.Get("param")

				if paramName == "" {
					continue
				}

				paramValue := r.PathValue(paramName)

				if field.CanSet() {
					field.SetString(paramValue)
				}
			}

			err := utils.Validate.Struct(payload)

			if err != nil {
				utils.WriteJsonErrorResponse(
					w,
					http.StatusBadRequest,
					"Validation failed.",
					err,
				)
				return
			}

			reqCtx := r.Context()
			ctx := context.WithValue(reqCtx, "payload", payload)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
