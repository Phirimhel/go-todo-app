package users_service

import (
	"context"

	"github.com/Phirimhel/go-todo-app/internal/core/domain"
	users_postgres_repository "github.com/Phirimhel/go-todo-app/internal/features/users/repository/postgres"
)

type UsersService interface {
	CreateUser(ctx context.Context, user domain.User) (domain.User, error)
	GetUsers(ctx context.Context, limit, offset *int) ([]domain.User, error)
	GetUser(ctx context.Context, id int) (domain.User, error)
	DeleteUser(ctx context.Context, id int) error
	PatchUser(ctx context.Context, id int, patch domain.UserPatch) (domain.User, error)
	LoginUser(ctx context.Context, userName, userPassword string) (string, error)
}

type userService struct {
	UsersRepository users_postgres_repository.UsersRepository
}

func NewUserService(repository users_postgres_repository.UsersRepository) *userService {
	return &userService{
		UsersRepository: repository,
	}
}
