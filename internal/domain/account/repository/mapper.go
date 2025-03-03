package repository

import (
	"komek/db/sqlc"
	"komek/internal/domain/account/entity"
	country "komek/internal/domain/country/entity"
	currency "komek/internal/domain/currency/entity"
)

func (r *Repository) mapAccount(acc sqlc.Account) entity.Account {
	return entity.Account{
		ID:          acc.ID.Bytes,
		Owner:       acc.Owner.Bytes,
		Balance:     acc.Balance,
		HoldBalance: acc.HoldBalance,
		Status:      entity.AccountStatus(acc.Status),
		Currency:    currency.Currency(acc.Currency),
		Country:     country.Country(acc.Country),
		CreatedAt:   acc.CreatedAt.Time,
		UpdatedAt:   acc.UpdatedAt.Time,
	}
}

func (r *Repository) mapAccounts(acc []sqlc.Account) entity.Accounts {
	res := make([]entity.Account, 0, len(acc))
	for _, a := range acc {
		res = append(res, r.mapAccount(a))
	}
	return res
}
