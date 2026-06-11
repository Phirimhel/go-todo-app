package tasks_service

import (
	"context"

	"github.com/Phirimhel/go-todo-app/internal/core/domain"
	tasks_repository "github.com/Phirimhel/go-todo-app/internal/features/tasks/repository/postgres"
)

type TasksService interface {
	CreateTask(ctx context.Context, task domain.Task) (domain.Task, error)
	GetTasks(ctx context.Context, authorID, limit, offset *int) ([]domain.Task, error)
	GetTask(ctx context.Context, taksID int) (domain.Task, error)
	DeleteTask(ctx context.Context, taksID int) error
}

type tasksService struct {
	repository tasks_repository.TasksRepository
}

func NewTasksService(repo tasks_repository.TasksRepository) *tasksService {
	return &tasksService{
		repository: repo,
	}
}
