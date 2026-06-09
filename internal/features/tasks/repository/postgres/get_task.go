package tasks_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/Phirimhel/go-todo-app/internal/core/domain"
	core_errors "github.com/Phirimhel/go-todo-app/internal/core/errors"
	core_postgres_pool "github.com/Phirimhel/go-todo-app/internal/core/repo/posgres/pool"
)

func (r *tasksRepository) GetTask(ctx context.Context, taskID int) (domain.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
		SELECT id, 
			version, 
			title, 
			description, 
			completed, 
			created_at, 
			completed_at,
			author_id
		FROM todoapp.tasks 
		WHERE (id = $1);
	`
	row := r.pool.QueryRow(ctx, query, taskID)

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
			"[repo]: scan model of selected tasks row: %w",
			err,
		)
	}

	taskDomain := taskDomainFromUserModel(taskModel)
	return taskDomain, nil

}
