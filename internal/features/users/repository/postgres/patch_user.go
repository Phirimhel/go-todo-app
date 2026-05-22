package users_postgres_repository

import (
	"context"
	"fmt"

	"github.com/Phirimhel/go-todo-app/internal/core/domain"
)

func (r *usersRepository) PatchUser(ctx context.Context, user domain.User) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
		UPDATE table_name
		SET full_name = $1, phone_number = $2
		WHERE id = $3;
	`

	row := r.pool.QueryRow(ctx, query, user.FullName, user.PhoneNumber, user.ID)

	var userModel UserModel
	if err := row.Scan(
		&userModel.ID,
		&userModel.Version,
		&userModel.FullName,
		&userModel.PhoneNumber,
	); err != nil {
		return domain.User{}, fmt.Errorf("[repo]: scan model: %w", err)
	}

	userDomen := userDomainFromUserModel(userModel)
	return userDomen, nil
}
