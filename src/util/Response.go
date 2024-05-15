package util

import (
	"encoding/json"
	"net/http"
)

func JsonError(w http.ResponseWriter, err error) {
	payload := struct {
		Message string `json:"message"`
	}{Message: err.Error()}

	JsonResponse(w, http.StatusBadRequest, payload)
}

func JsonResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		JsonError(w, err)
	}
}
