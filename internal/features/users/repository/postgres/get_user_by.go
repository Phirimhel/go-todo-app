package users_postgres_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/Phirimhel/go-todo-app/internal/core/domain"
	core_errors "github.com/Phirimhel/go-todo-app/internal/core/errors"
	core_postgres_pool "github.com/Phirimhel/go-todo-app/internal/core/repo/posgres/pool"
)

func (r *usersRepository) GetUserBy(ctx context.Context, email, fullName, phoneNumber *string) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
		SELECT id, Version, full_name, phone_number, role, email, password_hash FROM todoapp.users
		WHERE ($1::varchar IS NULL OR email = $1)
		  AND ($2::varchar IS NULL OR full_name = $2)
		  AND ($3::varchar IS NULL OR phone_number = $3);
	`

	row := r.pool.QueryRow(ctx, query, email, fullName, phoneNumber)

	var userModel UserModel
	if err := row.Scan(
		&userModel.ID,
		&userModel.Version,
		&userModel.FullName,
		&userModel.PhoneNumber,
		&userModel.Role,
		&userModel.Email,
		&userModel.PasswordHash,
	); err != nil {

		if errors.Is(err, core_postgres_pool.ErrNoRows) {
			return domain.User{}, fmt.Errorf(
				"[repo]: user not found, %w",
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
