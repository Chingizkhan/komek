package account_repo

import "errors"

const (
	ConstraintOwnerCurrencyKey        = "owner_currency_key"
	ConstraintAccountsOwnerForeignKey = "accounts_owner_fkey"
)

var (
	ErrAccountNotFound       = errors.New("account_not_found")
	ErrAccountAlreadyExists  = errors.New("account_already_exists")
	ErrCurrencyAlreadyExists = errors.New("currency_already_exists")
	ErrOwnerNotFound         = errors.New("owner_no_found")
)
