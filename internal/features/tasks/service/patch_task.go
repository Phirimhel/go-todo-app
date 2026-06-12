package tasks_service

import (
	"context"
	"fmt"

	"github.com/Phirimhel/go-todo-app/internal/core/domain"
)

func (s *tasksService) PatchTask(ctx context.Context, taskID int, taskPatch domain.TaskPatch) (domain.Task, error) {

	task, err := s.repository.GetTask(ctx, taskID)
	if err != nil {
		return domain.Task{}, fmt.Errorf("[service]: failed to get task %w", err)
	}

	if err := task.ApplyPatch(taskPatch); err != nil {
		return domain.Task{}, fmt.Errorf("[service]: faled to apply patch task: %w", err)
	}

	taskDomain, err := s.repository.PatchTask(ctx, taskID, task)
	if err != nil {
		return domain.Task{}, fmt.Errorf("[service]: failed to get task from repo %w", err)
	}

	return taskDomain, nil

}
