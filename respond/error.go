package respond

import (
	"fmt"
	"strconv"
)

type ErrorResponse struct {
	ErrorData errorResponseContent `json:"error"`
}

func (e ErrorResponse) Error() string {
	return fmt.Sprint(e.ErrorData)
}

func (e *ErrorResponse) AddValidation(validation errorResponseValidation) {
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

func SError(message string) ErrorResponse {
	return ErrorResponse{
		ErrorData: errorResponseContent{
			Message: message,
		},
	}
}

func SErrorFromErr(message string, err error) ErrorResponse {
	return ErrorResponse{
		ErrorData: errorResponseContent{
			Message: message,
			Details: err.Error(),
		},
	}
}

func Error(code int, message string, details string, validationErrors []errorResponseValidation) ErrorResponse {
	return ErrorResponse{
		ErrorData: errorResponseContent{
			strconv.Itoa(code),
			message,
			details,
			validationErrors,
		},
	}
}

func NewValidation(field string, message string) errorResponseValidation {
	return errorResponseValidation{
		Field:   field,
		Message: message,
	}
}
