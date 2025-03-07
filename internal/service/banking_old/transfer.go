package banking_old

//func (s *Banking) Transfer(ctx context.Context, in dto.TransferIn) (out dto.TransferOut, err error) {
//	response, err := s.client.Transfer(ctx, &pb.TransferIn{
//		FromAccountId: in.FromAccountID.String(),
//		ToAccountId:   in.ToAccountID.String(),
//		Amount:        in.Amount,
//	})
//	if err != nil {
//		return out, fmt.Errorf("transfer: %w", err)
//	}
//
//	transaction, err := mapper.ConvTransactionToDomain(response.Transaction)
//	if err != nil {
//		return out, fmt.Errorf("conv transaction: %w", err)
//	}
//
//	accFrom, err := mapper.ConvAccountProtoToDomain(response.AccountFrom)
//	if err != nil {
//		return out, fmt.Errorf("conv acc from: %w", err)
//	}
//
//	accTo, err := mapper.ConvAccountProtoToDomain(response.AccountTo)
//	if err != nil {
//		return out, fmt.Errorf("conv acc to: %w", err)
//	}
//
//	return dto.TransferOut{
//		Transaction: transaction,
//		FromAccount: accFrom,
//		ToAccount:   accTo,
//	}, nil
//}
