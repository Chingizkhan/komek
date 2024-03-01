package mapper

import (
	"komek/db/sqlc"
	"komek/internal/domain"
)

func ConvTransferToDomain(transfer sqlc.Transfer) domain.Transfer {
	return domain.Transfer{
		ID:            transfer.ID,
		FromAccountID: transfer.FromAccountID,
		ToAccountID:   transfer.ToAccountID,
		Amount:        transfer.Amount,
		CreatedAt:     transfer.CreatedAt.Time,
	}
}

func ConvEntryToDomain(entry sqlc.Entry) domain.Entry {
	return domain.Entry{
		ID:        entry.ID,
		AccountID: entry.AccountID,
		Amount:    entry.Amount,
		CreatedAt: entry.CreatedAt.Time,
	}
}

func ConvAccountToDomain(acc sqlc.Account) domain.Account {
	return domain.Account{
		ID:        acc.ID,
		Owner:     acc.Owner.Bytes,
		Balance:   acc.Balance,
		Currency:  acc.Currency,
		CreatedAt: acc.CreatedAt.Time,
	}
}
