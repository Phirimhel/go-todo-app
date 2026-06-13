package domain

import "time"

type Statistic struct {
	TasksCreated               int
	TasksCompleted             int
	TasksCompletedRate         *float64
	TasksAverageCompletionTime *time.Duration
}

func NewStatistic(
	tasksCreate int,
	tasksCompleted int,
	tasksCompletedRate *float64,
	tasksAverageCompletionTime *time.Duration,

) Statistic {
	return Statistic{
		TasksCreated:               tasksCreate,
		TasksCompleted:             tasksCompleted,
		TasksCompletedRate:         tasksCompletedRate,
		TasksAverageCompletionTime: tasksAverageCompletionTime,
	}
}
