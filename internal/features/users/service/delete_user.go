package users_service

import (
	"context"
	"fmt"
)

func (s *userService) DeleteUser(ctx context.Context, id int) error {
	if err := s.UsersRepository.DeleteUser(ctx, id); err != nil {
		return fmt.Errorf("[service]: failed to delete user with id='%d', %w", id, err)
	}
	return nil
}
