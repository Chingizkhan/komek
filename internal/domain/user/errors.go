package user

import "errors"

var (
	ErrInvalidLogin = errors.New("invalid login")
	ErrInvalidName  = errors.New("invalid name")
	ErrInvalidPhone = errors.New("invalid phone")
	ErrInvalidEmail = errors.New("invalid email")
)
