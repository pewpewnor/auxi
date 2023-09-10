package auxi

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/mitchellh/mapstructure"
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
		err := Respond.SError("No authorization header or value given")
		return "", err
	}

	values := strings.Split(value, " ")
	if len(values) != 2 {
		err := Respond.SError("Authorization header value is malformed")
		err.AddValidation(Respond.NewValidation(
			"Authorization header",
			"Expected exactly 2 values separated by spaces",
		))

		return "", err
	}
	if values[0] != tokenPrefix {
		err := Respond.SError("Authorization header value is malformed")
		err.AddValidation(Respond.NewValidation(
			"Authorization header",
			fmt.Sprintf("First value (token prefix) must be '%v'", tokenPrefix),
		))

		return "", err
	}

	return values[1], nil
}

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

func BindQueryString(r *http.Request, target any) error {
	targetReflectValue := reflect.ValueOf(target)
	if targetReflectValue.Kind() != reflect.Ptr && targetReflectValue.Elem().Kind() != reflect.Struct {
		panic("Target must be a pointer to a struct")
	}

	structType := targetReflectValue.Elem().Type()
	for i := 0; i < structType.NumField(); i++ {
		if field := structType.Field(i); field.Type.Kind() != reflect.String {
			panic(
				fmt.Sprintf(
					"Field '%s' of target struct is not of type string",
					field.Name))
		}
	}

	queryString := r.URL.Query()
	firstQueryString := map[string]string{}

	for key, values := range queryString {
		firstQueryString[key] = values[0]
	}

	return mapstructure.Decode(firstQueryString, target)
}
