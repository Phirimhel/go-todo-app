package users_postgres_repository

import (
	"context"
	"fmt"

	"github.com/Phirimhel/go-todo-app/internal/core/domain"
)

func (r *usersRepository) GetUsers(ctx context.Context, limit, offset *int) ([]domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
		SELECT id, Version, full_name, phone_number, role, email FROM todoapp.users
		ORDER BY id ASC
		LIMIT $1 
		OFFSET $2;
	`
	rows, err := r.pool.Query(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("[repo]: select users rows: %w", err)
	}
	defer rows.Close()

	var userModels []UserModel
	for rows.Next() {

		var userModel UserModel
		if err := rows.Scan(
			&userModel.ID,
			&userModel.Version,
			&userModel.FullName,
			&userModel.PhoneNumber,
			&userModel.Role,
			&userModel.Email,
		); err != nil {
			return nil, fmt.Errorf("[repo]: scan model of select users rows: %w", err)
		}

		userModels = append(userModels, userModel)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("next rows: %w", err)
	}

	userDomains := userDomainsFromUserModels(userModels)

	return userDomains, nil
}
