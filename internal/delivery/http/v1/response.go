package v1

import (
	"net/http"
	"encoding/json"
)

func jsonResponse(w http.ResponseWriter, statusCode int, data interface{}) {
    message, err := json.Marshal(data)
    w.Header().Set("Content-Type", "application/json")

    if err != nil {
        message, _ = json.Marshal(err.Error())
        w.WriteHeader(http.StatusInternalServerError)
        w.Write(message)
        return
    }

	w.WriteHeader(statusCode)
	w.Write(message)
	return
}