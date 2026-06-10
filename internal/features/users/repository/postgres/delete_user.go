package users_postgres_repository

import (
	"context"
	"fmt"

	core_errors "github.com/Phirimhel/go-todo-app/internal/core/errors"
)

func (r *usersRepository) DeleteUser(ctx context.Context, id int) error {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
		DELETE FROM todoapp.users 
		WHERE id = $1;
	`

	cmdTag, err := r.pool.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("[repo]: failed to delete user with id='%d' %w", id, err)
	}

	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("[repo]: user with id='%d' %w", id, core_errors.ErrNotFound)
	}

	return nil
}
