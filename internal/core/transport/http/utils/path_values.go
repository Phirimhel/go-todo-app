package core_http_utils

import (
	"fmt"
	"net/http"
	"strconv"

	core_errors "github.com/Phirimhel/go-todo-app/internal/core/errors"
)

func GetPathValue(r *http.Request, key string) (int, error) {
	pathValue := r.PathValue(key)

	if pathValue == "" {
		return 0, fmt.Errorf("[utils]: no keys=%s in path values. %w",
			key,
			core_errors.ErrInvalidArgument,
		)
	}

	val, err := strconv.Atoi(pathValue)
	if err != nil {
		return 0, fmt.Errorf("[utils]: failed to parse path value='%s' to int. %w",
			pathValue,
			core_errors.ErrInvalidArgument,
		)
	}

	return val, err
}
