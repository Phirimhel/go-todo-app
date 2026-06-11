package tasks_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/Phirimhel/go-todo-app/internal/core/domain"
	core_errors "github.com/Phirimhel/go-todo-app/internal/core/errors"
	core_postgres_pool "github.com/Phirimhel/go-todo-app/internal/core/repo/posgres/pool"
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
			completed_at,
			author_id
			) 
		VALUES ($1, $2, $3, $4, $5, $6)
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
		task.CompletedAt,
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

		if errors.Is(err, core_postgres_pool.ErrViolatesForeignKey) {
			return domain.Task{}, fmt.Errorf("[repo]: %v: failed to scan task model. author id = '%d' %w",
				err,
				task.AuthorID,
				core_errors.ErrNotFound,
			)
		}

		return domain.Task{}, fmt.Errorf("[repo]: failed to scan task model. %w", err)

	}

	taskDomen := taskDomainFromUserModel(TaskModel)

	return taskDomen, nil
}
