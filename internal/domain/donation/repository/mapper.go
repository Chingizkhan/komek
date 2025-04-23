package repository

import (
	"komek/db/sqlc"
	"komek/internal/domain/donation/entity"
)

func (r *Repository) mapDonation(in sqlc.GetDonationByIDRow) entity.Donation {
	return entity.Donation{
		ID:            in.ID.Bytes,
		FundraiseID:   in.FundraiseID.Bytes,
		TransactionID: in.TransactionID.Bytes,
		ClientID:      in.ClientID.Bytes,
		FromAccountID: in.FromAccountID.Bytes,
		ToAccountID:   in.ToAccountID.Bytes,
		Amount:        in.Amount,
		CreatedAt:     in.CreatedAt.Time.UTC(),
		ClientName:    in.ClientName.String,
		ClientPhoto:   in.ClientImage.String,
	}
}
