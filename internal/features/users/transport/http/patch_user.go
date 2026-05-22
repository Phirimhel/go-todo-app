package users_transport_http

import (
	"net/http"

	"github.com/Phirimhel/go-todo-app/internal/core/domain"
	core_logger "github.com/Phirimhel/go-todo-app/internal/core/logger"
	core_http_response "github.com/Phirimhel/go-todo-app/internal/core/transport/http/response"
)

type PatchUserRequest struct {
	FullName    *string `json:"full_name" validate:"required,min=3,max=100"`
	PhoneNumber *string `json:"phone_number" validate:"omitempty,e164"`
}

type PatchUserResponse UserDTOResponse

func (h *UsersHTTPHandler) PatchUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.GetLoggerFromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

	log.Debug("invoce patch user handler")

}

func domainFromPatchDTO(dto PatchUserRequest) domain.User {
	var userDomain domain.User
	if dto.FullName != nil {
		userDomain.FullName = *dto.FullName
	}

	if dto.PhoneNumber != nil {
		userDomain.FullName = *dto.PhoneNumber
	}
	return userDomain
}
