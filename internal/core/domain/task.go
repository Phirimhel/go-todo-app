package domain

import (
	"fmt"
	"time"

	core_errors "github.com/Phirimhel/go-todo-app/internal/core/errors"
)

type Task struct {
	ID      int
	Version int

	Title       string
	Description *string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time

	AuthorID int
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
			"invalid task title length: %d, %w",
			titleLength,
			core_errors.ErrInvalidArgument,
		)
	}

	if t.Description != nil {
		descriptionLength := len([]rune(*t.Description))
		if 1 > descriptionLength || descriptionLength > 1000 {
			return fmt.Errorf(
				"invalid task description length: %d, %w",
				descriptionLength,
				core_errors.ErrInvalidArgument,
			)
		}
	}

	// if created time < completed time
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

	// if completed but CompletedAt == nil
	if t.Completed && t.CompletedAt == nil {
		return fmt.Errorf(
			"'CompletedAt' can't be nil if field completed is 'true', %w",
			core_errors.ErrInvalidArgument,
		)
	}

	// if NOT completed but has CompletedAt NOT nil.
	if !t.Completed && t.CompletedAt != nil {
		return fmt.Errorf(
			"CompletedAt must be 'nil' if completed == false, %w",
			core_errors.ErrInvalidArgument,
		)
	}

	return nil
}

func (t *Task) ApplyPatch(patch TaskPatch) error {

	if err := patch.ValidatePatch(); err != nil {
		return fmt.Errorf("[task/domain]: failed validate task patch: %w", err)
	}

	tmp := *t
	if patch.Title.Set {
		tmp.Title = *patch.Title.Value
	}

	if patch.Description.Set {
		tmp.Description = patch.Description.Value
	}

	if patch.Completed.Set {

		if patch.Completed.Value == nil {
			return fmt.Errorf(
				"[task/domain]: faled to validate taks, copleted can't be null: %w",
				core_errors.ErrInvalidArgument,
			)
		}

		completed := *patch.Completed.Value
		tmp.Completed = completed

		if completed {
			completedAt := time.Now()
			tmp.CompletedAt = &completedAt
		} else {
			tmp.CompletedAt = nil
		}
	}

	if err := tmp.ValidateTask(); err != nil {
		return fmt.Errorf("[task/domain]: failed validate task: %w", err)
	}
	*t = tmp

	return nil
}

type TaskPatch struct {
	Title       Nullable[string] `json:"title"`
	Description Nullable[string] `json:"description"`
	Completed   Nullable[bool]   `json:"completed"`
}

func (p *TaskPatch) ValidatePatch() error {
	if p.Title.Set && p.Title.Value == nil {
		return fmt.Errorf(
			"title can't be patched to null %w",
			core_errors.ErrInvalidArgument,
		)
	}
	if p.Completed.Set && p.Completed.Value == nil {
		return fmt.Errorf(
			"completed can't be patched to null %w",
			core_errors.ErrInvalidArgument,
		)
	}

	return nil
}

func (t *Task) CompletionDuration() *time.Duration {
	if !t.Completed {
		return nil
	}

	if t.CompletedAt == nil {
		return nil
	}

	duration := t.CompletedAt.Sub(t.CreatedAt)

	return &duration
}
