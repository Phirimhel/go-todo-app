package users_service

import (
	"context"
	"fmt"

	"github.com/Phirimhel/go-todo-app/internal/core/domain"
)

func (u *userService) CreateUser(ctx context.Context, user domain.User) (domain.User, error) {

	// validate
	if err := user.Validation(); err != nil {
		return domain.User{}, fmt.Errorf("validate user domen: %w", err)
	}

	// repo
	user, err := u.UsersRepository.CreateUser(ctx, user)
	if err != nil {
		return domain.User{}, fmt.Errorf("create user repository: %w", err)
	}
	// return
	return user, nil
}
