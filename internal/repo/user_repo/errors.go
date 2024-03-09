package user_repo

import "errors"

var (
	ErrUserNotFound           = errors.New("user_not_found")
	ErrUserAlreadyExists      = errors.New("user_already_exists")
	ErrUserPhoneAlreadyExists = errors.New("phone_already_exists")
	ErrUserLoginAlreadyExists = errors.New("login_already_exists")
	ErrUserEmailAlreadyExists = errors.New("email_already_exists")
)

const (
	ConstraintUsersLoginKey = "users_login_key"
	ConstraintUsersPhoneKey = "users_phone_key"
	ConstraintUsersEmailKey = "users_email_key"
)
