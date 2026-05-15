package users_transport_http

import (
	"net/http"

	"github.com/Phirimhel/go-todo-app/internal/core/domain"
	core_logger "github.com/Phirimhel/go-todo-app/internal/core/logger"
	core_http_request "github.com/Phirimhel/go-todo-app/internal/core/transport/http/request"
	core_http_response "github.com/Phirimhel/go-todo-app/internal/core/transport/http/response"
)

type CreateUserRequest struct {
	FullName    string  `json:"full_name" validate:"required,min=3,max=100"`
	PhoneNumber *string `json:"phone_number" validate:"omitempty,e164"`
}

type CreateUserResponse struct {
	ID          int     `json:"id"`
	Version     int     `json:"version"`
	FullName    string  `json:"full_name"`
	PhoneNumber *string `json:"phone_number"`
}

func (u *UsersHTTPHandler) CreateUser(w http.ResponseWriter, r *http.Request) {

	log := core_logger.FromContext(r.Context())
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

	log.Debug("invoce create user handler")

	var request CreateUserRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode or validate request")
		return
	}

	newUserDomain := domainFromDTO(request)
	userDomain, err := u.usersService.CreateUser(r.Context(), newUserDomain)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to create user")
		return
	}

	response := dtoFromDomain(userDomain)
	responseHandler.JSONResponse(response, http.StatusCreated)
}

func domainFromDTO(dto CreateUserRequest) domain.User {
	return domain.NewUserUnitiliazied(dto.FullName, dto.PhoneNumber)
}

func dtoFromDomain(user domain.User) CreateUserResponse {
	return CreateUserResponse{
		ID:          user.ID,
		Version:     user.Version,
		FullName:    user.FullName,
		PhoneNumber: user.PhoneNumber,
	}

}
