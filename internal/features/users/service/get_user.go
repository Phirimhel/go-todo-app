package users_service

import (
	"context"
	"fmt"

	"github.com/Phirimhel/go-todo-app/internal/core/domain"
)

func (s *userService) GetUser(ctx context.Context, id int) (domain.User, error) {

	user, err := s.UsersRepository.GetUser(ctx, id)
	if err != nil {
		return domain.User{}, fmt.Errorf("[service]: failed to get user from repository: %w", err)
	}
	return user, err
}
