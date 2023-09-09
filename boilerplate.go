package auxi

import (
	"net/http"
	"strings"
)

func CORSMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.Header.Set("Access-Control-Allow-Origin", "*")
		r.Header.Set("Access-Control-Allow-Credentials", "true")
		r.Header.Set(
			"Access-Control-Allow-Headers",
			"Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		r.Header.Set(
			"Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		next.ServeHTTP(w, r)
	})
}

func GetBearerTokenFromAuthorizationHeader(r *http.Request) (string, error) {
	value := r.Header.Get("Authorization")
	if value == "" {
		err := Respond.SimpleError(
			"No authorization header or its value is not given")
		return "", err
	}

	values := strings.Split(value, " ")
	if len(values) != 2 {
		err := Respond.SimpleError("Authorization header value is malformed")
		err.AddValidation(Respond.CreateValidation(
			"Authorization header",
			"Expected exactly 2 values",
		))

		return "", err
	}
	if values[0] != "bearer" {
		err := Respond.SimpleError("Authorization header value is malformed")
		err.AddValidation(Respond.CreateValidation(
			"Authorization header",
			"First value must be 'Apikey'",
		))

		return "", err
	}

	return values[1], nil
}
