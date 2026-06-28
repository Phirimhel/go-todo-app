package core_auth

import (
	"fmt"

	core_errors "github.com/Phirimhel/go-todo-app/internal/core/errors"
	"github.com/golang-jwt/jwt/v5"
)

func (s *authenticator) ValidateJWT(tokenString string) (*Claims, error) {
	claims := Claims{}
	token, err := jwt.ParseWithClaims(tokenString, &claims, s.keyFunc)
	if err != nil {
		return nil, fmt.Errorf("[ValidateJWT]: jwt validation error: %w:, %w", err, core_errors.ErrUnauthorized)
	}

	if !token.Valid {
		return nil, fmt.Errorf("[ValidateJWT]: token is not valid: %w:, %w", err, core_errors.ErrUnauthorized)
	}

	_, err = token.Claims.GetSubject()
	if err != nil {
		return nil, fmt.Errorf("[ValidateJWT]: failed to get user id: %w:, %w", err, core_errors.ErrUnauthorized)
	}

	return &claims, nil
}
