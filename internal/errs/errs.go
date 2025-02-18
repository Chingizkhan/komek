package errs

import "errors"

var (
	ErrUserNotFound           = errors.New("user_not_found")
	ErrUserAlreadyExists      = errors.New("user_already_exists")
	ErrUserPhoneAlreadyExists = errors.New("phone_already_exists")
	ErrUserLoginAlreadyExists = errors.New("login_already_exists")
	ErrUserEmailAlreadyExists = errors.New("email_already_exists")

	IncorrectPassword = errors.New("incorrect_password")
)
