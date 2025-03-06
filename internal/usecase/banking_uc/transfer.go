package banking_uc

import (
	"context"
	"komek/internal/dto"
)

func (s *UseCase) Transfer(ctx context.Context, in dto.TransferIn) (out dto.TransferOut, err error) {
	panic("implement me")
	//out, err = s.banking.Transfer(ctx, in)
	//if err != nil {
	//	return out, fmt.Errorf("banking service -> transfer: %w", err)
	//}
	//return
}
