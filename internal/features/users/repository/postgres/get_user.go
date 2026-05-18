package users_postgres_repository

import (
	"context"

	"github.com/Phirimhel/go-todo-app/internal/core/domain"
)

func (u *usersRepository) GetUsers(ctx context.Context, limit, offset *int) ([]domain.User, error) {
	domain := make([]domain.User, 0)
	return domain, nil
}
