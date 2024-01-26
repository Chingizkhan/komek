package v1

import (
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"komek/internal/controller/http/api_util"
	UCWord "komek/internal/usecase/word"
	"komek/pkg/logger"
	"net/http"
)

type WordRemoveRequest struct {
	Value  string `json:"value"`
	UserID string `json:"user_id"`
}

type WordRemoveResponse struct {
	Status string `json:"status"`
}

func (h *Handler) remove(w http.ResponseWriter, r *http.Request) {
	const fnName = "word_http - remove"

	var req WordRemoveRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		h.l.Error(fnName, logger.Err(err))
		api_util.RenderErrorResponse(w, "invalid request", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	userID, err := uuid.Parse(req.UserID)
	if err != nil {
		h.l.Error(fnName, logger.Err(err))
		api_util.RenderErrorResponse(w, err.Error(), http.StatusConflict)
		return
	}

	err = h.wordUC.Delete(r.Context(), req.Value, userID)
	if err != nil {
		if errors.Is(err, UCWord.ErrNothingDeleted) {
			h.l.Error(fnName, logger.Err(err))
			api_util.RenderErrorResponse(w, err.Error(), http.StatusConflict)
			return
		}

		h.l.Error(fnName, logger.Err(err))
		api_util.RenderErrorResponse(w, "remove failed", http.StatusInternalServerError)
		return
	}

	api_util.RenderResponse(w,
		&WordRemoveResponse{
			Status: "success",
		},
		http.StatusCreated,
	)
}
