package http

import (
	"encoding/json"
	"net/http"
)

// ErrorResponse represents a response containing an error message.
type ErrorResponse struct {
	Error string `json:"error"`
}

func Err(w http.ResponseWriter, msg string, status int) {
	Resp(w, ErrorResponse{Error: msg}, status)
}

func Resp(w http.ResponseWriter, res interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")

	content, err := json.Marshal(res)
	if err != nil {
		// XXX Do something with the error ;)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(status)

	if _, err = w.Write(content); err != nil {
		// XXX Do something with the error ;)
	}
}
