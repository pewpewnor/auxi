package auxi

type SuccessResponseData struct {
	Status  string                 `json:"status"`
	Data    map[string]interface{} `json:"data"`
	Message string                 `json:"message"`
}

func SimpleSuccessResponse(message string) SuccessResponseData {
	return SuccessResponseData{
		Status:  "success",
		Message: message,
	}
}

func SuccessResponse(data map[string]interface{}, message string) SuccessResponseData {
	return SuccessResponseData{
		Status:  "success",
		Message: message,
		Data:    data,
	}
}
