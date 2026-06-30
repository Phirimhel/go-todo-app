package users_postgres_repository

import (
	"context"

	"github.com/Phirimhel/go-todo-app/internal/core/domain"
	core_postgres_pool "github.com/Phirimhel/go-todo-app/internal/core/repo/posgres/pool"
)

type UsersRepository interface {
	CreateUser(ctx context.Context, user domain.User) (domain.User, error)
	GetUsers(ctx context.Context, limit, offset *int) ([]domain.User, error)
	GetUser(ctx context.Context, id int) (domain.User, error)
	DeleteUser(ctx context.Context, id int) error
	PatchUser(ctx context.Context, id int, user domain.User) (domain.User, error)
	GetUserBy(ctx context.Context, email, fullName, phoneNumber *string) (domain.User, error)
}

type usersRepository struct {
	pool core_postgres_pool.Pool
}

func NewRepository(pool core_postgres_pool.Pool) *usersRepository {
	return &usersRepository{
		pool: pool,
	}
}
