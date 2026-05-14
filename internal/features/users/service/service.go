package users_service

import (
	"context"

	"github.com/Phirimhel/go-todo-app/internal/core/domain"
)

type UserService struct {
	UsersRepository UsersRepository
}

type UsersRepository interface {
	CreateUser(ctx context.Context, user domain.User) (domain.User, error)
}

func NewUserService(repository UsersRepository) *UserService {
	return &UserService{
		UsersRepository: repository,
	}
}
