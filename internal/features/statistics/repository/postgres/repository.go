package statistics_postgres_repository

import (
	"context"
	"time"

	"github.com/Phirimhel/go-todo-app/internal/core/domain"
	core_postgres_pool "github.com/Phirimhel/go-todo-app/internal/core/repo/posgres/pool"
)

type StatisticsRepository interface {
	GetTasks(ctx context.Context, userID *int, from, to *time.Time) ([]domain.Task, error)
}

type statisticsRepository struct {
	pool core_postgres_pool.Pool
}

func NewStatisticsRepository(pool core_postgres_pool.Pool) *statisticsRepository {
	return &statisticsRepository{
		pool: pool,
	}
}
