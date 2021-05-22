package v1

import (
	"net/http"
)

func jsonResponse(w http.ResponseWriter, statusCode int, message []byte) {
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	w.Write(message)
}
