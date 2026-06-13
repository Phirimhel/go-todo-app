package statistic_service

import (
	"context"
	"time"

	"github.com/Phirimhel/go-todo-app/internal/core/domain"
	statistics_postgres_repository "github.com/Phirimhel/go-todo-app/internal/features/statistics/repository/postgres"
)

type StatisticService interface {
	GetStatistics(ctx context.Context, userID *int, from, to *time.Time) (domain.Statistic, error)
}

type statisticService struct {
	repository statistics_postgres_repository.StatisticsRepository
}

func NewStatisticsSercice(repository statistics_postgres_repository.StatisticsRepository) *statisticService {
	return &statisticService{
		repository: repository,
	}
}
