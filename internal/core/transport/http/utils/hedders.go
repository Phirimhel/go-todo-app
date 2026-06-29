package core_http_utils

import (
	"fmt"
	"net/http"
	"strings"

	core_errors "github.com/Phirimhel/go-todo-app/internal/core/errors"
)

func GetBearerToken(headers http.Header) (string, error) {
	authHeader := headers.Get("Authorization")
	if authHeader == "" {
		return "", fmt.Errorf("missing Authorization header  %w", core_errors.ErrInvalidArgument)
	}
	token, found := strings.CutPrefix(authHeader, "Bearer ")
	if !found {
		return "", fmt.Errorf("invalid Authorization format %w", core_errors.ErrInvalidArgument)
	}

	return token, nil
}
