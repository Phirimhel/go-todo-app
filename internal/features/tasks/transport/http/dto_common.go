package tasks_transport_http

import (
	"time"

	"github.com/Phirimhel/go-todo-app/internal/core/domain"
)

type TaskDtoResponse struct {
	ID          int        `json:"id" example:"42"`
	Version     int        `json:"version" example:"1"`
	Title       string     `json:"title" example:"Buy groceries"`
	Description *string    `json:"description" example:"Milk, eggs, and bread"`
	Completed   bool       `json:"completed" example:"false"`
	CreatedAt   time.Time  `json:"created_at" example:"2026-06-20T12:00:00Z"`
	CompletedAt *time.Time `json:"completed_at" example:"2026-06-20T14:30:00Z"`
	AuthorID    int        `json:"author_id" example:"105"`
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
