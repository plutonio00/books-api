package v1

import (
	"encoding/json"
	"net/http"
)

func jsonResponse(w http.ResponseWriter, statusCode int, data interface{}) {
    var responseKey string

    if statusCode < http.StatusBadRequest {
        responseKey = "result"
    } else {
        responseKey = "error"
    }

    data = map[string]interface{}{responseKey: data}
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
