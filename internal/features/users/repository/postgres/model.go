package users_postgres_repository

import "github.com/Phirimhel/go-todo-app/internal/core/domain"

type UserModel struct {
	ID           int
	Version      int
	FullName     string
	PhoneNumber  *string
	Role         string
	Email        string
	PasswordHash string
}

func userDomainFromUserModel(model UserModel) domain.User {
	return domain.User{
		ID:           model.ID,
		Version:      model.Version,
		FullName:     model.FullName,
		PhoneNumber:  model.PhoneNumber,
		Role:         model.Role,
		Email:        model.Email,
		PasswordHash: model.PasswordHash,
	}

}

func userDomainsFromUserModels(models []UserModel) []domain.User {
	userDomains := make([]domain.User, len(models))

	for i, model := range models {
		userDomains[i] = userDomainFromUserModel(model)
	}

	return userDomains
}
