package users_postgres_repository

import (
	"context"
	"fmt"

	"github.com/Phirimhel/go-todo-app/internal/core/domain"
)

func (r *usersRepository) CreateUser(ctx context.Context, user domain.User) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
		INSERT INTO todoapp.users (full_name, phone_number, email) 
		VALUES ($1, $2, $3)
		RETURNING id, version, full_name, phone_number, email;
	`

	row := r.pool.QueryRow(ctx, query, user.FullName, user.PhoneNumber, user.Email)

	var userModel UserModel
	if err := row.Scan(
		&userModel.ID,
		&userModel.Version,
		&userModel.FullName,
		&userModel.PhoneNumber,
		&userModel.Role,
		&userModel.Email,
	); err != nil {
		return domain.User{}, fmt.Errorf("[repo]: scan model: %w", err)
	}

	userDomen := userDomainFromUserModel(userModel)
	return userDomen, nil
}
