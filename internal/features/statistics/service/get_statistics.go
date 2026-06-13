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

	_ = tasksDomains

	return domain.Statistic{}, nil
}

func calcStatistics(tasks []domain.Task) domain.Statistic {

	if len(tasks) == 0 {
		return domain.Statistic{}
	}

	tasksCreated := len(tasks)

	tasksCompleted := 0
	for i := range tasks {
		if tasks[i].Completed {
			tasksCompleted++
		}
	}

	tasksCompletedRate := float64(tasksCreated) / float64(tasksCompleted) * 100

	//tasksAverageCompletionTime := time.Duration

	return domain.Statistic{
		TasksCreated:       tasksCreated,
		TasksCompleted:     tasksCompleted,
		TasksCompletedRate: &tasksCompletedRate,
		//TasksAverageCompletionTime: &tasksAverageCompletionTime,
	}
}
