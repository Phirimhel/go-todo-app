package users_service

import (
	"context"
	"fmt"

	"github.com/Phirimhel/go-todo-app/internal/core/domain"
)

func (s *userService) CreateUser(ctx context.Context, user domain.User) (domain.User, error) {

	if err := user.ValidateUser(); err != nil {
		return domain.User{}, fmt.Errorf("[service]: validate user domen: %w", err)
	}

	user, err := s.UsersRepository.CreateUser(ctx, user)
	if err != nil {
		return domain.User{}, fmt.Errorf("[service]: create user in repository: %w", err)
	}

	return user, nil
}
