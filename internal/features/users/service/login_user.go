package users_service

import (
	"context"
	"fmt"

	core_auth "github.com/Phirimhel/go-todo-app/internal/core/auth"
	"github.com/Phirimhel/go-todo-app/internal/core/domain"
	core_errors "github.com/Phirimhel/go-todo-app/internal/core/errors"
)

func (s *userService) LoginUser(ctx context.Context, email, pasword string) (domain.User, error) {

	userDomain, err := s.UsersRepository.GetUserBy(ctx, &email, nil, nil)
	if err != nil {
		return domain.User{}, fmt.Errorf("[service]: failed to get user from repository: %w", err)
	}

	isValid, err := core_auth.CheckPasswordHash(pasword, userDomain.PasswordHash)
	if err != nil {
		return domain.User{}, fmt.Errorf("[service]: failed to check user's password: %w", core_errors.ErrUnauthorized)
	}

	if !isValid {
		return domain.User{}, fmt.Errorf("[service]: user pasword is not walid: %w", core_errors.ErrUnauthorized)
	}

	return userDomain, nil
}
