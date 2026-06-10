package tasks_transport_http

import (
	"time"

	"github.com/Phirimhel/go-todo-app/internal/core/domain"
)

type TaskDtoResponse struct {
	ID          int        `json:"id"`
	Version     int        `json:"version"`
	Title       string     `json:"title"`
	Description *string    `json:"description"`
	Completed   bool       `json:"completed"`
	CreatedAt   time.Time  `json:"created_at"`
	CompletedAt *time.Time `json:"completed_at"`
	AuthorID    int        `json:"author_id"`
}

func taskDTOfromDomain(task domain.Task) TaskDtoResponse {
	return TaskDtoResponse{
		ID:          task.ID,
		Version:     task.Version,
		Title:       task.Title,
		Description: task.Description,
		Completed:   task.Completed,
		CreatedAt:   task.CreatedAt,
		CompletedAt: task.CompletedAt,
		AuthorID:    task.AuthorID,
	}
}

func tasksDTOfromDomain(tasks []domain.Task) []TaskDtoResponse {
	tasksDTO := make([]TaskDtoResponse, len(tasks))
	for i := range tasks {
		tasksDTO[i] = taskDTOfromDomain(tasks[i])
	}
	return tasksDTO
}
