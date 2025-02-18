package user

import "errors"

var (
	ErrSessionBlocked       = errors.New("session_is_blocked")
	ErrSessionUser          = errors.New("incorrect_session_user")
	ErrMismatchSessionToken = errors.New("mismatch_session_token")
	ErrExpiredSession       = errors.New("expired_session")
)
