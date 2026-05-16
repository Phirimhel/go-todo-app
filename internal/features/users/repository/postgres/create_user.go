package users_postgres_repository

import (
	"context"
	"fmt"

	"github.com/Phirimhel/go-todo-app/internal/core/domain"
)

func (r *UsersRepository) CreateUser(ctx context.Context, user domain.User) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
		INSERT INTO todoapp.users (full_name, phone_number) 
		VALUES ($1, $2)
		RETURNING id, version, full_name, phone_number;
	`

	if r.pool == nil {
		panic("users repository pool is nil")
	}

	row := r.pool.QueryRow(ctx, query, user.FullName, user.PhoneNumber)

	var model UserModel
	if err := row.Scan(
		&model.ID,
		&model.Version,
		&model.FullName,
		&model.PhoneNumber,
	); err != nil {
		return domain.User{}, fmt.Errorf("scan model: %w", err)
	}

	createdUser := domain.User{
		ID:          model.ID,
		Version:     model.Version,
		FullName:    model.FullName,
		PhoneNumber: model.PhoneNumber,
	}

	return createdUser, nil
}
