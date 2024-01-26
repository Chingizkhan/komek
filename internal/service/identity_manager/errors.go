package identity_manager

import "errors"

var (
	ErrUserNotFound = errors.New("user not found")
	ErrNoSession    = errors.New("user doesn't have session")
)
