package tasks_service

import (
	"context"
	"fmt"

	"github.com/Phirimhel/go-todo-app/internal/core/domain"
)

func (s *tasksService) GetTask(ctx context.Context, taskID int) (domain.Task, error) {

	taskDomain, err := s.repository.GetTask(ctx, taskID)
	if err != nil {
		return domain.Task{}, fmt.Errorf(
			"[service]: failed to get task by id: '%d', %w",
			taskID,
			err,
		)
	}
	return taskDomain, nil
}
