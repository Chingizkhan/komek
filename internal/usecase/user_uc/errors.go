package user_uc

import "errors"

var (
	ErrIncorrectPassword    = errors.New("incorrect_password")
	ErrSessionBlocked       = errors.New("session_is_blocked")
	ErrSessionUser          = errors.New("incorrect_session_user")
	ErrMismatchSessionToken = errors.New("mismatch_session_token")
	ErrExpiredSession       = errors.New("expired_session")
)
