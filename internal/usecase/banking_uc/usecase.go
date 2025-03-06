package banking_uc

import (
	"komek/internal/usecase"
)

type (
	UseCase struct {
		tr      usecase.Transactional
		account usecase.AccountService
		banking usecase.BankingService
	}
)

var txKey = struct{}{}

func New(
	tr usecase.Transactional,
	banking usecase.BankingService,
	account usecase.AccountService,
) *UseCase {
	return &UseCase{
		tr:      tr,
		account: account,
		banking: banking,
	}
}
