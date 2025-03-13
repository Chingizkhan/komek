package repository

import (
	"komek/db/sqlc"
	"komek/internal/domain/fundraise/entity"
)

func (r *Repository) mapFundraises(fundraises []sqlc.Fundraise) []entity.Fundraise {
	result := make([]entity.Fundraise, len(fundraises))
	for i, f := range fundraises {
		result[i] = r.mapFundraise(f)
	}
	return result
}

func (r *Repository) mapFundraise(f sqlc.Fundraise) entity.Fundraise {
	return entity.Fundraise{
		ID:        f.ID.Bytes,
		Goal:      f.Goal,
		Collected: f.Collected,
		AccountID: f.AccountID.Bytes,
		IsActive:  f.IsActive.Bool,
	}
}
