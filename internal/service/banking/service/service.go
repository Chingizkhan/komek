package service

type Service struct {
	transaction TransactionService
	operation   OperationService
	account     AccountService
}

func New(
	operation OperationService,
	transaction TransactionService,
	account AccountService,
) *Service {
	return &Service{
		operation:   operation,
		transaction: transaction,
		account:     account,
	}
}
