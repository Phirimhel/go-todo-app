package statistic_service

import (
	"context"
	"fmt"
	"time"

	"github.com/Phirimhel/go-todo-app/internal/core/domain"
	core_errors "github.com/Phirimhel/go-todo-app/internal/core/errors"
)

func (s *statisticService) GetStatistics(ctx context.Context, userID *int, from, to *time.Time) (domain.Statistic, error) {

	if from != nil && to != nil {
		if from.After(*to) || from.Equal(*to) {
			return domain.Statistic{}, fmt.Errorf("[service]: to must be after from %w", core_errors.ErrInvalidArgument)
		}
	}

	tasksDomains, err := s.repository.GetTasks(ctx, userID, from, to)
	if err != nil {
		return domain.Statistic{}, fmt.Errorf("[service]: failed to get statistic domain %w", err)
	}

	statisticsDomain := calcStatistics(tasksDomains)

	return statisticsDomain, nil
}

func calcStatistics(tasks []domain.Task) domain.Statistic {

	if len(tasks) == 0 {
		return domain.Statistic{}
	}

	// total tasks
	tasksCreated := len(tasks)

	// total tasks completed
	tasksCompleted := 0
	var totalCompletedDuration time.Duration
	for i := range tasks {
		if tasks[i].Completed {
			tasksCompleted++

		}

		completedDuration := tasks[i].CompletionDuration()
		if completedDuration != nil {
			totalCompletedDuration += *completedDuration
		}

	}

	// rate of completed tasks
	tasksCompletedRate := float64(tasksCompleted) / float64(tasksCreated) * 100.0

	// avg of completion time per task
	var tasksAverageCompletionTime *time.Duration
	if tasksCompleted > 0 && totalCompletedDuration != 0 {

		avg := totalCompletedDuration / time.Duration(tasksCompleted)
		tasksAverageCompletionTime = &avg
	}

	return domain.Statistic{
		TasksCreated:               tasksCreated,
		TasksCompleted:             tasksCompleted,
		TasksCompletedRate:         &tasksCompletedRate,
		TasksAverageCompletionTime: tasksAverageCompletionTime,
	}
}
