package tasks_postgres_repository

import (
	"context"

	"github.com/Phirimhel/go-todo-app/internal/core/domain"
	core_postgres_pool "github.com/Phirimhel/go-todo-app/internal/core/repo/posgres/pool"
)

type TasksRepository interface {
	CreateTask(ctx context.Context, task domain.Task) (domain.Task, error)
	GetTasks(ctx context.Context, id, limit, offset *int) ([]domain.Task, error)
	GetTask(ctx context.Context, taskID int) (domain.Task, error)
	DeleteTask(ctx context.Context, taksID int) error
	PatchTask(ctx context.Context, id int, user domain.Task) (domain.Task, error)
}

type tasksRepository struct {
	pool core_postgres_pool.Pool
}

func NewTasksRepository(pool core_postgres_pool.Pool) *tasksRepository {
	return &tasksRepository{
		pool: pool,
	}
}
