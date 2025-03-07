package errs

import "errors"

var (
	ErrUserNotFound           = errors.New("user_not_found")
	ErrUserAlreadyExists      = errors.New("user_already_exists")
	ErrUserPhoneAlreadyExists = errors.New("phone_already_exists")
	ErrUserLoginAlreadyExists = errors.New("login_already_exists")
	ErrUserEmailAlreadyExists = errors.New("email_already_exists")
	UserGetParamsNotSpecified = errors.New("user_get_params_not_specified")

	ErrClientNotFound = errors.New("client_not_found")

	IncorrectPassword = errors.New("incorrect_password")

	// account
	AccountNotFound       = errors.New("account_not_found")
	AccountStatusInvalid  = errors.New("invalid_account_status")
	InactiveAccountStatus = errors.New("inactive_account_status")
	FromAccountNotFound   = errors.New("from_account_not_found")
	ToAccountNotFound     = errors.New("to_account_not_found")

	CurrencyMismatch = errors.New("currency_mismatch")

	// operation
	OperationNotFound = errors.New("operation_not_found")

	// transaction
	TransactionNotFound = errors.New("transaction_not_found")

	InsufficientAmount = errors.New("insufficient_amount")
	NotEnoughBalance   = errors.New("not_enough_balance")
)
