package hasher

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type PasswordHash struct{}

func NewPasswordHash() *PasswordHash {
	return &PasswordHash{}
}

func (p *PasswordHash) HashPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	return string(hashedBytes), nil
}

func (p *PasswordHash) ComparePassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
