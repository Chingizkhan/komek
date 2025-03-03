package repository

import (
	"komek/db/sqlc"
	"komek/internal/domain/account/entity"
	"komek/pkg/null_value"
)

type CreateAccountRequest struct {
	in entity.CreateIn
}

func (m CreateAccountRequest) toSqlc() sqlc.CreateAccountParams {
	return sqlc.CreateAccountParams{
		Owner:    null_value.UUID(m.in.Owner),
		Balance:  m.in.Balance,
		Country:  string(m.in.Country),
		Currency: string(m.in.Currency),
	}
}
