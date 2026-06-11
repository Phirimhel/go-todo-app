package tasks_service

import (
	"context"
	"fmt"
)

func (s *tasksService) DeleteTask(ctx context.Context, taskID int) error {
	if err := s.repository.DeleteTask(ctx, taskID); err != nil {
		return fmt.Errorf(
			"[service]: failed to delete task with id: %d, %w",
			taskID,
			err,
		)
	}
	return nil
}
