package v1

import (
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"komek/internal/controller/http/api_util"
	"komek/internal/domain/word"
	"komek/pkg/logger"
	"net/http"
)

type WordGetRequest struct {
	Value  string    `json:"value"`
	UserID uuid.UUID `json:"user_id"`
}

func (r *WordGetRequest) fill(req *http.Request) error {
	value, _ := mux.Vars(req)["word"]
	r.Value = value

	userID, _ := mux.Vars(req)["userID"]
	userId, err := uuid.Parse(userID)
	if err != nil {
		return err
	}
	r.UserID = userId

	return nil
}

type WordGetResponse struct {
	Body word.Word
}

func (h *Handler) get(w http.ResponseWriter, r *http.Request) {
	const fnName = "word_http - get"

	var req WordGetRequest
	err := req.fill(r)
	if err != nil {
		h.l.Error(fnName, logger.Err(err))
		api_util.RenderErrorResponse(w, err.Error(), http.StatusConflict)
		return
	}

	wordModel, err := h.wordUC.Get(r.Context(), req.Value, req.UserID)
	if err != nil {
		h.l.Error(fnName, logger.Err(err))
		api_util.RenderErrorResponse(w, "get failed", http.StatusInternalServerError)
		return
	}

	resp := &WordGetResponse{
		Body: wordModel,
	}

	api_util.RenderResponse(w,
		resp.Body,
		http.StatusCreated,
	)
}
