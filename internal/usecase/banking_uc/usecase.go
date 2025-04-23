package banking_uc

import (
	"komek/internal/usecase"
)

type (
	UseCase struct {
		tr          usecase.Transactional
		account     usecase.AccountService
		banking     usecase.BankingService
		funds       usecase.FundraiseService
		transaction usecase.TransactionService
		donation    usecase.DonationService
	}
)

var txKey = struct{}{}

func New(
	tr usecase.Transactional,
	banking usecase.BankingService,
	account usecase.AccountService,
	funds usecase.FundraiseService,
	transaction usecase.TransactionService,
	donation usecase.DonationService,
) *UseCase {
	return &UseCase{
		tr:          tr,
		account:     account,
		banking:     banking,
		funds:       funds,
		transaction: transaction,
		donation:    donation,
	}
}
