package auxi

import (
	"strconv"
)

type ErrorResponseData struct {
	ErrorData errorResponseContent `json:"error"`
}

func (e *ErrorResponseData) AddValidation(validation ErrorResponseValidation) {
	e.ErrorData.ValidationErrors = append(e.ErrorData.ValidationErrors,
		validation,
	)
}

func (e ErrorResponseData) Error() string {
	return e.ErrorData.Message
}

type errorResponseContent struct {
	Code             string                    `json:"code"`
	Message          string                    `json:"message"`
	Details          string                    `json:"details"`
	ValidationErrors []ErrorResponseValidation `json:"validationErrors"`
}

type ErrorResponseValidation struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func SimpleErrorResponse(message string, details string) ErrorResponseData {
	return ErrorResponseData{
		ErrorData: errorResponseContent{
			Message: message,
			Details: details,
		},
	}
}

func SimpleErrorResponseFromError(message string, err error) ErrorResponseData {
	return ErrorResponseData{
		ErrorData: errorResponseContent{
			Message: message,
			Details: err.Error(),
		},
	}
}

func ErrorResponse(code int, message string, details string, validationErrors []ErrorResponseValidation) ErrorResponseData {
	return ErrorResponseData{
		ErrorData: errorResponseContent{
			strconv.Itoa(code),
			message,
			details,
			validationErrors,
		},
	}
}
