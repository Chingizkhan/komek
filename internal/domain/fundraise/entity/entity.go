package entity

import "github.com/google/uuid"

type (
	Fundraise struct {
		ID        uuid.UUID `json:"id"`
		Goal      int64     `json:"goal"`
		Collected int64     `json:"collected"`
		Type      Type      `json:"type"`
		AccountID uuid.UUID `json:"account_id"`
		IsActive  bool      `json:"is_active"`
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
)
