package statistics_transport_http

import (
	"github.com/Phirimhel/go-todo-app/internal/core/domain"
)

type StatisticDTOResponse struct {
	TasksCreate                int      `json:"tasks_created" example:"15"`
	TasksCompleted             int      `json:"tasks_completed" example:"12"`
	TasksCompletedRate         *float64 `json:"tasks_completed_rate" example:"80.0"`
	TasksAverageCompletionTime *string  `json:"tasks_average_completed_time" example:"2h 15m"`
}

func StatisticDTOfromDomain(statistics domain.Statistic) StatisticDTOResponse {

	var avgTime *string

	if statistics.TasksAverageCompletionTime != nil {
		duration := statistics.TasksAverageCompletionTime.String()
		avgTime = &duration
	}

	return StatisticDTOResponse{
		TasksCreate:                statistics.TasksCreated,
		TasksCompleted:             statistics.TasksCompleted,
		TasksCompletedRate:         statistics.TasksCompletedRate,
		TasksAverageCompletionTime: avgTime,
	}
}
