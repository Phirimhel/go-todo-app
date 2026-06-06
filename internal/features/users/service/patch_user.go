package users_service

import (
	"context"
	"fmt"

	"github.com/Phirimhel/go-todo-app/internal/core/domain"
)

func (s *userService) PatchUser(ctx context.Context, id int, patch domain.UserPatch) (domain.User, error) {

	user, err := s.GetUser(ctx, id)
	if err != nil {
		return domain.User{}, fmt.Errorf("[service]: failed to get user with id: %d,  %w", id, err)
	}

	if err := user.ApplyPatch(patch); err != nil {
		return domain.User{}, fmt.Errorf("[service]: faled to patch user: %w", err)
	}
	userDomain, err := s.UsersRepository.PatchUser(ctx, id, user)
	if err != nil {
		return domain.User{}, fmt.Errorf("[service]: faled to patch user: %w", err)
	}

	return userDomain, nil
}
