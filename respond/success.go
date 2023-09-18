package respond

type SuccessResponseData struct {
	Status  string         `json:"status"`
	Data    map[string]any `json:"data"`
	Message string         `json:"message"`
}

func SSuccess(message string) SuccessResponseData {
	return SuccessResponseData{
		Status:  "success",
		Message: message,
	}
}

func Success(data map[string]any, message string) SuccessResponseData {
	return SuccessResponseData{
		Status:  "success",
		Message: message,
		Data:    data,
	}
}
