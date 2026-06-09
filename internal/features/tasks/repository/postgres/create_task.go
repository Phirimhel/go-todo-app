package tasks_repository

import (
	"context"
	"fmt"

	"github.com/Phirimhel/go-todo-app/internal/core/domain"
)

func (r *tasksRepository) CreateTask(ctx context.Context, task domain.Task) (domain.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
		INSERT INTO todoapp.tasks (
			title, 
			description, 
			completed, 
			created_at, 
			author_id
			) 
		VALUES ($1, $2, $3, $4, $5)
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

	row := r.pool.QueryRow(
		ctx,
		query,
		task.Title,
		task.Description,
		task.Completed,
		task.CreatedAt,
		task.AuthorID,
	)

	var TaskModel TaskModel
	if err := row.Scan(
		&TaskModel.ID,
		&TaskModel.Version,
		&TaskModel.Title,
		&TaskModel.Description,
		&TaskModel.Completed,
		&TaskModel.CreatedAt,
		&TaskModel.CompletedAt,
		&TaskModel.AuthorID,
	); err != nil {
		return domain.Task{}, fmt.Errorf("[repo]: failed to scan task model %w", err)
	}

	taskDomen := taskDomainFromUserModel(TaskModel)

	return taskDomen, nil
}
