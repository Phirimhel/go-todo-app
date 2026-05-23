package users_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/Phirimhel/go-todo-app/internal/core/domain"
	core_errors "github.com/Phirimhel/go-todo-app/internal/core/errors"
	"github.com/jackc/pgx/v5"
)

func (r *usersRepository) GetUser(ctx context.Context, id int) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
	SELECT id, Version, full_name, phone_number FROM todoapp.users
	WHERE id = $1 
	`

	row := r.pool.QueryRow(ctx, query, id)

	var userModel UserModel
	if err := row.Scan(
		&userModel.ID,
		&userModel.Version,
		&userModel.FullName,
		&userModel.PhoneNumber,
	); err != nil {

		if errors.Is(err, pgx.ErrNoRows) {
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
