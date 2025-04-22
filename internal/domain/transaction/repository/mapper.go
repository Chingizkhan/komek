package repository

import (
	"komek/db/sqlc"
	"komek/internal/domain/transaction/entity"
)

func (r *Repository) transactionsToDomain(transactions []sqlc.Transaction) []entity.Transaction {
	result := make([]entity.Transaction, 0, len(transactions))
	for _, transaction := range transactions {
		result = append(result, r.transactionToDomain(transaction))
	}
	return result
}

func (r *Repository) transactionToDomain(transaction sqlc.Transaction) entity.Transaction {
	return entity.Transaction{
		ID:            transaction.ID.Bytes,
		FromAccountID: transaction.FromAccountID.Bytes,
		ToAccountID:   transaction.ToAccountID.Bytes,
		Amount:        transaction.Amount,
		CreatedAt:     transaction.CreatedAt.Time,
	}
}

func (r *Repository) donationsToDomain(donations []sqlc.GetDonationsByAccountsRow) []entity.Donation {
	result := make([]entity.Donation, 0, len(donations))
	for _, donation := range donations {
		result = append(result, r.donationToDomain(donation))
	}
	return result
}

func (r *Repository) donationToDomain(donation sqlc.GetDonationsByAccountsRow) entity.Donation {
	return entity.Donation{
		ID:            donation.ID.Bytes,
		FromAccountID: donation.FromAccountID.Bytes,
		ToAccountID:   donation.ToAccountID.Bytes,
		Amount:        donation.Amount,
		CreatedAt:     donation.CreatedAt.Time,
		ClientName:    donation.ClientName.String,
		ClientPhoto:   donation.ClientPhoto.String,
	}
}
