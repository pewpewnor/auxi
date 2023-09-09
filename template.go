package auxi

import (
	"fmt"
	"net/http"
	"strings"
)

type authorizationHeader struct {
	r *http.Request
}

func NewAuthorizationHeader(r *http.Request) authorizationHeader {
	return authorizationHeader{
		r: r,
	}
}

func (ah authorizationHeader) GetBearerToken(tokenPrefix string) (string, error) {
	value := ah.r.Header.Get("Authorization")
	if value == "" {
		err := Respond.SimpleError("No authorization header or value given")
		return "", err
	}

	values := strings.Split(value, " ")
	if len(values) != 2 {
		err := Respond.SimpleError("Authorization header value is malformed")
		err.AddValidation(Respond.CreateValidation(
			"Authorization header",
			"Expected exactly 2 values separated by spaces",
		))

		return "", err
	}
	if values[0] != tokenPrefix {
		err := Respond.SimpleError("Authorization header value is malformed")
		err.AddValidation(Respond.CreateValidation(
			"Authorization header",
			fmt.Sprintf("First value (token prefix) must be '%v'", tokenPrefix),
		))

		return "", err
	}

	return values[1], nil
}

type Middleware func(next http.HandlerFunc) http.HandlerFunc

func NewCORSMiddleware(options map[string]string) Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			r.Header.Set("Access-Control-Allow-Origin", "*")
			r.Header.Set("Access-Control-Allow-Credentials", "true")
			r.Header.Set(
				"Access-Control-Allow-Headers",
				"Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
			r.Header.Set(
				"Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

			for k, v := range options {
				r.Header.Set(k, v)
			}

			next.ServeHTTP(w, r)
		})
	}
}
