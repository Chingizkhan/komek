package dto

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"komek/internal/domain"
	"net/http"
	"strconv"
)

type (
	TransferIn struct {
		FromAccountID int64
		ToAccountID   int64
		Amount        int64
	}

	TransferOut struct {
		Transfer    domain.Transfer
		FromAccount domain.Account
		ToAccount   domain.Account
		FromEntry   domain.Entry
		ToEntry     domain.Entry
	}

	CreateAccountIn struct {
		Owner    string
		Balance  int64
		Currency string
	}

	GetAccountIn struct {
		ID int64
	}
)

func (in *CreateAccountIn) ParseAndValidate(r *http.Request) error {
	err := json.NewDecoder(r.Body).Decode(in)
	if err != nil {
		return fmt.Errorf("decode response body: %w", err)
	}
	return nil
}

func (in *TransferIn) ParseAndValidate(r *http.Request) error {
	err := json.NewDecoder(r.Body).Decode(in)
	if err != nil {
		return fmt.Errorf("decode response body: %w", err)
	}
	defer r.Body.Close()
	return nil
}

func (in *GetAccountIn) ParseAndValidate(r *http.Request) error {
	idString := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		return fmt.Errorf("strconv.Atoi: %w", err)
	}
	in.ID = int64(id)
	return nil
}
