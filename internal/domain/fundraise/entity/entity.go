package entity

import "github.com/google/uuid"

type (
	Fundraise struct {
		ID        uuid.UUID
		Goal      int64
		Collected int64
		AccountID uuid.UUID
		IsActive  bool
	}

	CreateIn struct {
		Goal      int64
		Collected int64
		AccountID uuid.UUID
		IsActive  *bool
	}
)
