package users_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/Phirimhel/go-todo-app/internal/core/domain"
	core_errors "github.com/Phirimhel/go-todo-app/internal/core/errors"
	core_postgres_pool "github.com/Phirimhel/go-todo-app/internal/core/repo/posgres/pool"
)

func (r *usersRepository) GetUser(ctx context.Context, id int) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	SELECT id, Version, full_name, phone_number, role, email FROM todoapp.users
	WHERE id = $1; 
	`

	row := r.pool.QueryRow(ctx, query, id)

	var userModel UserModel
	if err := row.Scan(
		&userModel.ID,
		&userModel.Version,
		&userModel.FullName,
		&userModel.PhoneNumber,
		&userModel.Role,
		&userModel.Email,
	); err != nil {

		if errors.Is(err, core_postgres_pool.ErrNoRows) {
			return domain.User{}, fmt.Errorf(
				"[repo]: user with id='%d' %w",
				id,
				core_errors.ErrNotFound,
			)
		}

		return domain.User{}, fmt.Errorf(
			"[repo]: scan model of selected users row: %w",
			err,
		)
	}

	userDomen := userDomainFromUserModel(userModel)

	return userDomen, nil
}
