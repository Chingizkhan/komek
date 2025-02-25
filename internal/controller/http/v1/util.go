package v1

import (
	"encoding/json"
	customMiddleware "komek/internal/controller/http/middleware"
	"komek/internal/service/token"
	"komek/pkg/logger"
	"net/http"
)

// ErrorResponse represents a response containing an error message.
type ErrorResponse struct {
	Error string `json:"error"`
}

func (h *Handler) Err(w http.ResponseWriter, msg string, status int) {
	h.Resp(w, ErrorResponse{Error: msg}, status)
}

func (h *Handler) Error(w http.ResponseWriter, err error, status int) {
	h.l.Error("error http", logger.Err(err))
	err = h.Unwrap(err)
	h.Resp(w, ErrorResponse{Error: err.Error()}, status)
}

func (h *Handler) Resp(w http.ResponseWriter, res interface{}, status int) {
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

func (h *Handler) RespAnother(w http.ResponseWriter, res interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		//w.WriteHeader(http.StatusInternalServerError) // write header can be used once?
		http.Error(w, "internal_error", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) Unwrap(err error) error {
	type wrapper interface {
		Unwrap() error
	}

	for err != nil {
		wrapped, ok := err.(wrapper)
		if !ok {
			break
		}
		err = wrapped.Unwrap()
	}
	return err
}

func (h *Handler) payload(r *http.Request) *token.Payload {
	return r.Context().Value(customMiddleware.AuthorizationPayloadKey).(*token.Payload)
}
