package tasks_postgres_repository

import (
	"context"
	"fmt"

	core_errors "github.com/Phirimhel/go-todo-app/internal/core/errors"
)

func (r *tasksRepository) DeleteTask(ctx context.Context, taskID int) error {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
		DELETE FROM todoapp.tasks 
		WHERE id = $1;
	`

	comandTag, err := r.pool.Exec(ctx, query, taskID)

	if err != nil {
		return fmt.Errorf("[repo]: falide to delete task, %w", err)
	}

	if comandTag.RowsAffected() == 0 {
		return fmt.Errorf("[repo]: falide to delete task, %w", core_errors.ErrNotFound)
	}

	return nil
}
