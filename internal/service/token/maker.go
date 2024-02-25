package token

import (
	"github.com/google/uuid"
	"time"
)

type Maker interface {
	CreateToken(userID uuid.UUID, duration time.Duration) (string, error)
	VerifyToken(token string) (payload *Payload, err error)
}
