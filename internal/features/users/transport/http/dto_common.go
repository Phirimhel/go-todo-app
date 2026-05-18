package users_transport_http

import "github.com/Phirimhel/go-todo-app/internal/core/domain"

type UserDTOResponse struct {
	ID          int     `json:"id"`
	Version     int     `json:"version"`
	FullName    string  `json:"full_name"`
	PhoneNumber *string `json:"phone_number"`
}

func userDTOFromDomain(user domain.User) UserDTOResponse {
	return UserDTOResponse{
		ID:          user.ID,
		Version:     user.Version,
		FullName:    user.FullName,
		PhoneNumber: user.PhoneNumber,
	}

}

func usersDTOFromDomains(userDomains []domain.User) []UserDTOResponse {
	usersDTO := make([]UserDTOResponse, len(userDomains))

	for i := range userDomains {
		usersDTO[i] = userDTOFromDomain(userDomains[i])
	}

	return usersDTO

}
