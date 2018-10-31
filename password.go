package util

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword ...
func HashPassword(pass string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return hash, nil
}

// ValidPassword ...
func ValidPassword(pass string, hashedPassword []byte) error {
	if err := bcrypt.CompareHashAndPassword(hashedPassword, []byte(pass)); err != nil {
		return err
	}
	return nil
}
