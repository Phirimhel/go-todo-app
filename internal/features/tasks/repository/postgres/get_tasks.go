package tasks_postgres_repository

import (
	"context"
	"fmt"

	"github.com/Phirimhel/go-todo-app/internal/core/domain"
)

func (r *tasksRepository) GetTasks(ctx context.Context, authorID, limit, offset *int) ([]domain.Task, error) {
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
		WHERE ($1::int IS NULL OR author_id = $1)
		ORDER BY id ASC
		LIMIT $2 
		OFFSET $3;
	`

	args := []any{authorID, limit, offset}

	rows, err := r.pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("[repo]: select tasks rows: %w", err)
	}
	defer rows.Close()

	var tasksModels []TaskModel
	for rows.Next() {

		var taskModel TaskModel
		if err := rows.Scan(
			&taskModel.ID,
			&taskModel.Version,
			&taskModel.Title,
			&taskModel.Description,
			&taskModel.Completed,
			&taskModel.CreatedAt,
			&taskModel.CompletedAt,
			&taskModel.AuthorID,
		); err != nil {
			return nil, fmt.Errorf("[repo]: scan task row: %w", err)
		}

		tasksModels = append(tasksModels, taskModel)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("[repo]: rows error: %w", err)
	}

	taskDomains := tasksDomainsFromUserModels(tasksModels)

	return taskDomains, nil
}
