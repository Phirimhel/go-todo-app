package tasks_service

import (
	"context"
	"fmt"

	"github.com/Phirimhel/go-todo-app/internal/core/domain"
)

func (s *tasksService) CreateTask(ctx context.Context, task domain.Task) (domain.Task, error) {

	if err := task.ValidateTask(); err != nil {
		return domain.Task{}, fmt.Errorf("[service]: validate task domen %w", err)
	}

	task, err := s.repository.CreateTask(ctx, task)
	if err != nil {
		return domain.Task{}, fmt.Errorf("[service]: failed to create user in repo %w", err)
	}

	return task, nil
}
