package tasks_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/Phirimhel/go-todo-app/internal/core/domain"
	core_errors "github.com/Phirimhel/go-todo-app/internal/core/errors"
	core_postgres_pool "github.com/Phirimhel/go-todo-app/internal/core/repo/posgres/pool"
)

func (r *tasksRepository) PatchTask(ctx context.Context, taskID int, task domain.Task) (domain.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
		UPDATE todoapp.tasks
		SET 
			title = $1, 
			description = $2, 
			completed = $3,
			completed_at = $4,
			version = version+1
		WHERE id = $5 AND version = $6
		RETURNING
			id,
			version, 
			title, 
			description, 
			completed, 
			created_at, 
			completed_at,
			author_id;
	`

	params := []any{
		task.Title,
		task.Description,
		task.Completed,
		task.CompletedAt,
		taskID,
		task.Version,
	}

	row := r.pool.QueryRow(ctx, query, params...)

	var taskModel TaskModel
	if err := row.Scan(
		&taskModel.ID,
		&taskModel.Version,
		&taskModel.Title,
		&taskModel.Description,
		&taskModel.Completed,
		&taskModel.CreatedAt,
		&taskModel.CompletedAt,
		&taskModel.AuthorID,
	); err != nil {
		if errors.Is(err, core_postgres_pool.ErrNoRows) {
			return domain.Task{}, fmt.Errorf(
				"[repo]: task with id='%d' %w",
				taskID,
				core_errors.ErrNotFound,
			)
		}

		return domain.Task{}, fmt.Errorf(
			"[repo]: scan into model error: %w",
			err,
		)
	}

	taskDomain := taskDomainFromUserModel(taskModel)

	return taskDomain, nil
}
