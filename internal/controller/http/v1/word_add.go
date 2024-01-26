package v1

import (
	"encoding/json"
	"komek/internal/controller/http/api_util"
	"komek/internal/domain/word"
	"komek/pkg/logger"
	"net/http"
)

type WordAddRequest struct {
	Value       string `json:"value"`
	Language    string `json:"language"`
	Translation string `json:"translation"`
}

type WordAddResponse struct {
	Status string `json:"status"`
}

func (h *Handler) add(w http.ResponseWriter, r *http.Request) {
	const fnName = "word_http - add"

	var req WordAddRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.l.Error(fnName, logger.Err(err))
		api_util.RenderErrorResponse(w, "invalid request", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	err := h.wordUC.Save(r.Context(), word.Word{
		Value:       req.Value,
		Language:    word.Language(req.Language),
		Translation: req.Translation,
	})
	if err != nil {
		h.l.Error(fnName, logger.Err(err))
		api_util.RenderErrorResponse(w, "add failed", http.StatusInternalServerError)
		return
	}

	api_util.RenderResponse(w,
		&WordAddResponse{Status: "success"},
		http.StatusCreated,
	)
}
