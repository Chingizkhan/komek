package hasher

import (
	"golang.org/x/crypto/bcrypt"
)

type Hasher struct {
}

func New() *Hasher {
	return &Hasher{}
}

func (h *Hasher) Hash(value string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(value), 14)
	return string(bytes), err
}

func (h *Hasher) CheckHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
