package banking_uc

import (
	"komek/internal/usecase"
)

type (
	UseCase struct {
		tr      usecase.Transactional
		account usecase.AccountService
		banking usecase.BankingService
		funds   usecase.FundraiseService
	}
)

var txKey = struct{}{}

func New(
	tr usecase.Transactional,
	banking usecase.BankingService,
	account usecase.AccountService,
	funds usecase.FundraiseService,
) *UseCase {
	return &UseCase{
		tr:      tr,
		account: account,
		banking: banking,
		funds:   funds,
	}
}
