package auxi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

var Respond respondWith

type ErrorResponseData struct {
	ErrorData errorResponseContent `json:"error"`
}

func (e ErrorResponseData) Error() string {
	return e.ErrorData.Message
}

func (e *ErrorResponseData) AddValidation(validation errorResponseValidation) {
	e.ErrorData.ValidationErrors = append(e.ErrorData.ValidationErrors,
		validation)
}

type errorResponseValidation struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type errorResponseContent struct {
	Code             string                    `json:"code"`
	Message          string                    `json:"message"`
	Details          string                    `json:"details"`
	ValidationErrors []errorResponseValidation `json:"validationErrors"`
}

type SuccessResponseData struct {
	Status  string         `json:"status"`
	Data    map[string]any `json:"data"`
	Message string         `json:"message"`
}

type respondWith struct{}

func (rw respondWith) JSON(w http.ResponseWriter, data any, httpStatusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")

	response, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		defer panic(fmt.Sprintf(
			"Could not marshal response %v to JSON: %v", response, err))

		errorResponse, err := json.Marshal(rw.SimpleErrorFromErr(
			"Could not marshal response to JSON", err))
		if err != nil {
			w.Write([]byte("Could not marshal response to JSON"))
		}

		w.Write(errorResponse)
		return
	}

	w.WriteHeader(httpStatusCode)
	w.Write(response)
}

func (rw respondWith) SimpleError(message string) ErrorResponseData {
	return ErrorResponseData{
		ErrorData: errorResponseContent{
			Message: message,
		},
	}
}

func (rw respondWith) SimpleErrorFromErr(message string, err error) ErrorResponseData {
	return ErrorResponseData{
		ErrorData: errorResponseContent{
			Message: message,
			Details: err.Error(),
		},
	}
}

func (rw respondWith) Error(code int, message string, details string, validationErrors []errorResponseValidation) ErrorResponseData {
	return ErrorResponseData{
		ErrorData: errorResponseContent{
			strconv.Itoa(code),
			message,
			details,
			validationErrors,
		},
	}
}

func (rw respondWith) CreateValidation(field string, message string) errorResponseValidation {
	return errorResponseValidation{
		Field:   field,
		Message: message,
	}
}

func (rw respondWith) SimpleSuccess(message string) SuccessResponseData {
	return SuccessResponseData{
		Status:  "success",
		Message: message,
	}
}

func (rw respondWith) Success(data map[string]any, message string) SuccessResponseData {
	return SuccessResponseData{
		Status:  "success",
		Message: message,
		Data:    data,
	}
}
