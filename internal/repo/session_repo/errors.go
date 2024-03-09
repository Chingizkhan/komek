package session_repo

import "errors"

var (
	ErrSessionNotFound       = errors.New("session_not_found")
	ErrSessionUserIDNotFound = errors.New("user_id_not_found")
)

const (
	ConstraintUserIDFKey = "sessions_user_id_fkey"
)
