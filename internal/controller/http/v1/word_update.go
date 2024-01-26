package v1

import (
	"encoding/json"
	"errors"
	"komek/internal/controller/http/api_util"
	"komek/internal/domain/word"
	UCWord "komek/internal/usecase/word"
	"komek/pkg/logger"
	"net/http"
)

type WordUpdateRequest struct {
	OldValue    string `json:"old_value"`
	Value       string `json:"value"`
	Language    string `json:"language"`
	Translation string `json:"translation"`
}

type WordUpdateResponse struct {
	Status string `json:"status"`
}

func (h *Handler) update(w http.ResponseWriter, r *http.Request) {
	const fnName = "word_http - update"

	var req WordUpdateRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		h.l.Error(fnName, logger.Err(err))
		api_util.RenderErrorResponse(w, "invalid request", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	_, err = h.wordUC.Update(r.Context(), req.OldValue, word.Word{
		Value:       req.Value,
		Language:    word.Language(req.Language),
		Translation: req.Translation,
	})
	if err != nil {
		if errors.Is(err, UCWord.ErrNothingUpdated) {
			h.l.Error(fnName, logger.Err(err))
			api_util.RenderErrorResponse(w, err.Error(), http.StatusInternalServerError)
			return
		}

		h.l.Error(fnName, logger.Err(err))
		api_util.RenderErrorResponse(w, "update failed", http.StatusInternalServerError)
		return
	}

	api_util.RenderResponse(w,
		&WordUpdateResponse{
			Status: "success",
		},
		http.StatusCreated,
	)
}
