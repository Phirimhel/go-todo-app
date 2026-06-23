package web_fs_repository

import (
	"errors"
	"fmt"
	"os"

	core_errors "github.com/Phirimhel/go-todo-app/internal/core/errors"
)

func (r *webRepository) GetFile(path string) ([]byte, error) {

	file, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, fmt.Errorf("[repo]: file with path: '%s' is not found %w",
				path,
				core_errors.ErrNotFound,
			)
		}

		return nil, fmt.Errorf("[repo]: failed to get file '%s' %w",
			path,
			err,
		)

	}
	return file, nil
}
