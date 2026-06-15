package statistics_postgres_repository

import (
	"time"

	"github.com/Phirimhel/go-todo-app/internal/core/domain"
)

type TaskModel struct {
	ID          int
	Version     int
	Title       string
	Description *string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
	AuthorID    int
}

func taskDomainFromUserModel(model TaskModel) domain.Task {
	return domain.Task{
		ID:          model.ID,
		Version:     model.Version,
		Title:       model.Title,
		Description: model.Description,
		Completed:   model.Completed,
		CreatedAt:   model.CreatedAt,
		CompletedAt: model.CompletedAt,
		AuthorID:    model.AuthorID,
	}

}

func tasksDomainsFromUserModels(models []TaskModel) []domain.Task {
	domains := make([]domain.Task, len(models))

	for i := range domains {
		domains[i] = taskDomainFromUserModel(models[i])
	}

	return domains
}
