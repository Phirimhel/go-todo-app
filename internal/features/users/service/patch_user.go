package users_service

import (
	"context"
	"fmt"

	"github.com/Phirimhel/go-todo-app/internal/core/domain"
)

func (s *userService) PatchUser(ctx context.Context, user domain.User) (domain.User, error) {

	if err := user.Validation(); err != nil {
		return domain.User{}, fmt.Errorf("[service]: validate user domen: %w", err)
	}

	user, err := s.UsersRepository.PatchUser(ctx, user)
	if err != nil {
		return domain.User{}, fmt.Errorf("[service]: create user repository: %w", err)
	}

	return user, nil
}
