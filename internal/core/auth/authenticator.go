package core_auth

import (
	"context"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

type claimsContextKey struct{}
type authenticatorContextKey struct{}

var ClaimsContextKey = claimsContextKey{}
var AuthenticatorContextKey = authenticatorContextKey{}

type Authenticator interface {
	MakeJWT(userID, role string) (tokenString string, err error)
	ValidateJWT(tokenString string) (*Claims, error)
}

type authenticator struct {
	config Config
}

func MustGetClaimsFromContext(ctx context.Context) *Claims {
	claims, ok := ctx.Value(ClaimsContextKey).(*Claims)
	if !ok {
		panic("missing jwt claims in context")
	}
	return claims
}

func (s *authenticator) keyFunc(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}
	return []byte(s.config.JWTSecret), nil
}

func NewAuthenticator(config Config) Authenticator {
	return &authenticator{
		config: config,
	}
}
