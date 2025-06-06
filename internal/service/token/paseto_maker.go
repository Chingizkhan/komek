package token

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/o1egl/paseto"
	"golang.org/x/crypto/chacha20poly1305"
	"time"
)

type PasetoMaker struct {
	paseto       *paseto.V2
	symmetricKey []byte
}

func NewPasetoMaker(symmetricKey string) (Maker, error) {
	if len(symmetricKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid key size: must be exactly %d characters", chacha20poly1305.KeySize)
	}

	return &PasetoMaker{
		paseto:       paseto.NewV2(),
		symmetricKey: []byte(symmetricKey),
	}, nil
}

func (maker PasetoMaker) CreateToken(userID uuid.UUID, duration time.Duration) (string, *Payload, error) {
	payload, err := NewPayload(userID, duration)
	if err != nil {
		return "", nil, fmt.Errorf("NewPayload: %w", err)
	}

	token, err := maker.paseto.Encrypt(maker.symmetricKey, payload, nil)
	if err != nil {
		return "", nil, fmt.Errorf("paseto.Encrypt: %w", err)
	}

	return token, payload, nil
}

func (maker PasetoMaker) VerifyToken(token string) (payload *Payload, err error) {
	payload = new(Payload)
	if err = maker.paseto.Decrypt(token, maker.symmetricKey, payload, nil); err != nil {
		return nil, ErrInvalidToken
	}
	if err = payload.Valid(); err != nil {
		return nil, err
	}
	return payload, nil
}
