package users_postgres_repository

import core_postgres_pool "github.com/Phirimhel/go-todo-app/internal/core/repo/posgres/pool"

type UsersRepository struct {
	pool core_postgres_pool.Pool
}

type Repository interface{}

func NewRepository(pool core_postgres_pool.Pool) *UsersRepository {
	return &UsersRepository{
		pool: pool,
	}
}
