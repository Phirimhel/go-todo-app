package users_transport_http

import "github.com/Phirimhel/go-todo-app/internal/core/domain"

type UserDTOResponse struct {
	ID          int     `json:"id" example:"10"`
	Version     int     `json:"version" example:"23"`
	FullName    string  `json:"full_name" example:"John Doe"`
	PhoneNumber *string `json:"phone_number" example:"+35921234567"`
	Role        string  `json:"role" example:"admin"`
}

func userDTOFromDomain(user domain.User) UserDTOResponse {
	return UserDTOResponse{
		ID:          user.ID,
		Version:     user.Version,
		FullName:    user.FullName,
		PhoneNumber: user.PhoneNumber,
		Role:        user.Role,
	}
}

func usersDTOFromDomains(userDomains []domain.User) []UserDTOResponse {
	usersDTO := make([]UserDTOResponse, len(userDomains))

	for i := range userDomains {
		usersDTO[i] = userDTOFromDomain(userDomains[i])
	}

	return usersDTO

}
