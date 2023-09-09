package hres

import (
	"strconv"
)

type ErrorResponseValidation struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type errorResponseContent struct {
	Code             string                    `json:"code"`
	Message          string                    `json:"message"`
	Details          string                    `json:"details"`
	ValidationErrors []ErrorResponseValidation `json:"validationErrors"`
}

type ErrorResponseData struct {
	ErrorData errorResponseContent `json:"error"`
}

func (e ErrorResponseData) Error() string {
	return e.ErrorData.Message
}

func (e *ErrorResponseData) AddValidation(validation ErrorResponseValidation) {
	e.ErrorData.ValidationErrors = append(
		e.ErrorData.ValidationErrors, validation)
}

func SimpleError(message string) ErrorResponseData {
	return ErrorResponseData{
		ErrorData: errorResponseContent{
			Message: message,
		},
	}
}

func SimpleErrorFromErr(message string, err error) ErrorResponseData {
	return ErrorResponseData{
		ErrorData: errorResponseContent{
			Message: message,
			Details: err.Error(),
		},
	}
}

func ValidationError(code int, message string, details string, validationErrors []ErrorResponseValidation) ErrorResponseData {
	return ErrorResponseData{
		ErrorData: errorResponseContent{
			strconv.Itoa(code),
			message,
			details,
			validationErrors,
		},
	}
}
