package tasks_service

import (
	"context"
	"fmt"

	"github.com/Phirimhel/go-todo-app/internal/core/domain"
	core_errors "github.com/Phirimhel/go-todo-app/internal/core/errors"
)

func (s *tasksService) GetTasks(ctx context.Context, authorID, limit, offset *int) ([]domain.Task, error) {
	if limit != nil && *limit < 0 {
		return nil, fmt.Errorf("[service]: limit must be non-negative: %w", core_errors.ErrInvalidArgument)
	}

	if offset != nil && *offset < 0 {
		return nil, fmt.Errorf("[service]: limit must be non-negative: %w", core_errors.ErrInvalidArgument)
	}

	userDomains, err := s.repository.GetTasks(ctx, authorID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("[service]: failed to get tasks from repository %w", err)
	}

	return userDomains, nil
}
