package hasher

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type Hasher struct {
}

func New() *Hasher {
	return &Hasher{}
}

func (h *Hasher) Hash(value string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(value), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	return string(bytes), nil
}

func (h *Hasher) CheckHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
