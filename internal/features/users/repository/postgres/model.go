package users_postgres_repository

import "github.com/Phirimhel/go-todo-app/internal/core/domain"

type UserModel struct {
	ID          int
	Version     int
	FullName    string
	PhoneNumber *string
}

func userDomainFromUserModel(userModel []UserModel) []domain.User {
	userDomains := make([]domain.User, len(userModel))

	for i, model := range userModel {

		userDomains[i] = domain.User{
			ID:          model.ID,
			Version:     model.Version,
			FullName:    model.FullName,
			PhoneNumber: model.PhoneNumber,
		}
	}

	return userDomains
}
