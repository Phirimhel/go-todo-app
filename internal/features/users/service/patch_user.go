package users_service

import (
	"context"

	"github.com/Phirimhel/go-todo-app/internal/core/domain"
)

func (s *userService) PatchUser(ctx context.Context, id int, patch domain.UserPatch) (domain.User, error) {

	// 1. get user by id
	// 2. apply patch to user
	// 3. save patched user in repository

	return domain.User{}, nil
}
