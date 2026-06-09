package domain

import (
	"fmt"
	"time"

	core_errors "github.com/Phirimhel/go-todo-app/internal/core/errors"
)

type Task struct {
	ID          int
	Version     int
	Title       string
	Description *string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
	AuthorID    int
}

func NewTaskUnitialized(
	title string,
	description *string,
	completed bool,
	authorId int,
) Task {
	return NewTask(
		UnitiliaziedID,
		UnitiliaziedVersion,
		title,
		description,
		completed,
		time.Now(),
		nil,
		authorId,
	)
}

func NewTask(
	id, version int,
	title string,
	description *string,
	completed bool,
	createdAt time.Time,
	completedAt *time.Time,
	authorId int,
) Task {
	return Task{
		ID:          id,
		Version:     version,
		Title:       title,
		Description: description,
		Completed:   completed,
		CreatedAt:   createdAt,
		CompletedAt: completedAt,
		AuthorID:    authorId,
	}
}

func (t *Task) ValidateTask() error {

	titleLength := len([]rune(t.Title))
	if titleLength < 1 || titleLength > 100 {
		return fmt.Errorf(
			"invalid FullName length: %d, %w",
			titleLength,
			core_errors.ErrInvalidArgument,
		)
	}

	if t.Description != nil {
		descriptionLength := len([]rune(*t.Description))
		if 1 > descriptionLength || descriptionLength > 1000 {
			return fmt.Errorf(
				"invalid Description length: %d, %w",
				descriptionLength,
				core_errors.ErrInvalidArgument,
			)
		}
	}

	if t.CompletedAt != nil {
		if t.CreatedAt.After(*t.CompletedAt) {
			return fmt.Errorf(
				"task can't be completed (%v) before it is created (%v), %w",
				*t.CompletedAt,
				t.CreatedAt,
				core_errors.ErrInvalidArgument,
			)
		}
	}

	return nil
}
