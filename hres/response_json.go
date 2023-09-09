package hres

import (
	"encoding/json"
	"net/http"
)

func RespondWithJSON(w http.ResponseWriter, data any, httpStatusCode int) {
	response, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Server could not encode response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(httpStatusCode)
	w.Write(response)
}
