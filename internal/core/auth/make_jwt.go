package core_auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	jwt.RegisteredClaims
	Role string `json:"role"`
}

func (c *Claims) GetRole() (string, error) {
	return c.Role, nil
}

func (s *authenticator) MakeJWT(userID string, role string) (string, error) {

	claims := s.NewClaimsWithRole(userID, role)
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	secret := []byte(s.config.JWTSecret)

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", fmt.Errorf("[MakeJWT]: failed to signe toekn %W", err)
	}

	return tokenString, nil
}

func (s *authenticator) NewClaimsWithRole(userID, role string) Claims {
	claimsWIthRole := Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    s.config.Issuer,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.config.TokenTTL)),
			Subject:   userID,
		},
		Role: role,
	}

	return claimsWIthRole
}
