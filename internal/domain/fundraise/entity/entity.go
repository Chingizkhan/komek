package entity

import (
	"github.com/google/uuid"
)

type (
	Fundraise struct {
		ID                 uuid.UUID `json:"id"`
		Goal               int64     `json:"goal"`
		Collected          int64     `json:"collected"`
		Type               Type      `json:"type"`
		AccountID          uuid.UUID `json:"account_id"`
		IsActive           bool      `json:"is_active"`
		SupportersQuantity int64     `json:"supporters_quantity"`
	}

	Type struct {
		ID   uuid.UUID
		Name string
	}

	CreateIn struct {
		Goal      int64
		Collected int64
		TypeID    uuid.UUID
		AccountID uuid.UUID
		IsActive  *bool
	}

	ListOut struct {
		ID         uuid.UUID `json:"id"`
		Name       string    `json:"name"`
		ImageUrl   string    `json:"image_url"`
		City       string    `json:"city"`
		Categories []string  `json:"categories"`
		Goal       int64     `json:"goal"`
		Collected  int64     `json:"collected"`
	}

	GetOut struct {
		ID                 uuid.UUID `json:"id"`
		Name               string    `json:"name"`
		ImageUrl           string    `json:"image_url"`
		City               string    `json:"city"`
		Categories         []string  `json:"categories"`
		Goal               int64     `json:"goal"`
		Collected          int64     `json:"collected"`
		Description        string    `json:"description"`
		SupportersQuantity int64     `json:"supporters_quantity"`
		AccountID          uuid.UUID `json:"account_id"`
		IsActive           bool      `json:"is_active"`
	}
)
