package mapper

import (
	"komek/db/sqlc"
	"komek/internal/domain"
)

func ConvSessionToDomain(session sqlc.Session) domain.Session {
	return domain.Session{
		ID:           session.ID.Bytes,
		UserID:       session.UserID.Bytes,
		RefreshToken: session.RefreshToken,
		UserAgent:    session.UserAgent,
		ClientIp:     session.ClientIp,
		IsBlocked:    session.IsBlocked,
		ExpiresAt:    session.ExpiresAt.Time,
		CreatedAt:    session.CreatedAt.Time,
	}
}
