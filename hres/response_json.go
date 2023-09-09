package hres

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func RespondWithJSON(w http.ResponseWriter, data any, httpStatusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")

	response, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		defer panic(fmt.Sprintf("Could not marshal response %v to JSON: %v", response, err))

		errorResponse, err := json.Marshal(SimpleErrorFromErr("Could not marshal response to JSON", err))
		if err != nil {
			w.Write([]byte("Could not marshal response to JSON"))
		}

		w.Write(errorResponse)
		return
	}

	w.WriteHeader(httpStatusCode)
	w.Write(response)
}
