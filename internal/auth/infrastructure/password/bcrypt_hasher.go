package password

import (
	"golang.org/x/crypto/bcrypt"
)

type Hasher interface {
	Hash(password string) (string, error)
	Verify(hashedPassword, password string) bool
}

type bcryptHasher struct {
	cost int
}

func NewBcryptHasher() Hasher {
	return &bcryptHasher{cost: 12}
}

func (h *bcryptHasher) Hash(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), h.cost)
	return string(hashedPassword), err
}

func (h *bcryptHasher) Verify(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
