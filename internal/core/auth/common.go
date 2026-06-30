package core_auth

import (
	"fmt"

	"github.com/alexedwards/argon2id"
)

func HashPassword(password string) (string, error) {
	hash, err := argon2id.CreateHash(password, argon2id.DefaultParams)
	if err != nil {
		return "", fmt.Errorf("[hash pasword]: faled to has pasword: %w", err)
	}

	return hash, nil
}

func CheckPasswordHash(password, hash string) (bool, error) {

	isValid, err := argon2id.ComparePasswordAndHash(password, hash)
	if err != nil {
		return false, fmt.Errorf("[hash pasword]: faled to validate hash: %w", err)
	}

	return isValid, nil
}
