package entity

import (
	"github.com/google/uuid"
	"komek/internal/domain/email"
	"komek/internal/domain/phone"
	"time"
)

type (
	Client struct {
		ID            uuid.UUID   `json:"id"`
		Name          string      `json:"name"`
		Phone         phone.Phone `json:"phone"`
		Email         email.Email `json:"email"`
		Age           int         `json:"age"`
		City          string      `json:"city"`
		Address       string      `json:"address"`
		Description   string      `json:"description"`
		Circumstances string      `json:"circumstances"`
		ImageURL      string      `json:"image_url"`
		Categories    Categories  `json:"categories"`
		CreatedAt     time.Time   `json:"created_at"`
		UpdatedAt     time.Time   `json:"updated_at"`
	}

	Category struct {
		ID   uuid.UUID `json:"id"`
		Name string    `json:"name"`
	}
	Categories []Category

	Clients []Client

	CreateIn struct {
		Name          string      `json:"name"`
		Phone         phone.Phone `json:"phone"`
		Email         email.Email `json:"email"`
		Age           int         `json:"age"`
		City          string      `json:"city"`
		Address       string      `json:"address"`
		Description   string      `json:"description"`
		Circumstances string      `json:"circumstances"`
		ImageURL      string      `json:"image_url"`
		CategoryIDs   []uuid.UUID `json:"category_ids"`
	}

	BindCategories struct {
		ClientID    uuid.UUID   `json:"client_id"`
		CategoryIDs []uuid.UUID `json:"category_ids"`
	}
)

func (c Categories) Names() []string {
	res := make([]string, len(c))
	for i, category := range c {
		res[i] = category.Name
	}
	return res
}

func (c Categories) IDs() []uuid.UUID {
	res := make([]uuid.UUID, len(c))
	for i, category := range c {
		res[i] = category.ID
	}
	return res
}
