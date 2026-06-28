package users_transport_http

import (
	"net/http"

	"github.com/Phirimhel/go-todo-app/internal/core/domain"
	core_logger "github.com/Phirimhel/go-todo-app/internal/core/logger"
	core_http_request "github.com/Phirimhel/go-todo-app/internal/core/transport/http/request"
	core_http_response "github.com/Phirimhel/go-todo-app/internal/core/transport/http/response"
)

// @name CreateUserRequest
type CreateUserRequest struct {
	FullName    string  `json:"full_name" validate:"required,min=3,max=100" example:"John Doe"`
	PhoneNumber *string `json:"phone_number" validate:"omitempty,e164" example:"+35921234567"`
}

type CreateUserResponse UserDTOResponse

// CreateUser godoc
// @Summary      Create user
// @Description  Create new user
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        request   body      CreateUserRequest  true  "Create user request body"
// @Success      201       {object}  CreateUserResponse  "User successfully created"
// @Failure      400       {object}  core_http_response.ErrorResponse  "Bad request"
// @Failure      500       {object}  core_http_response.ErrorResponse  "Internal server error"
// @Router       /users [post]
func (h *UsersHTTPHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.MustGetLoggerFromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

	log.Debug("invoce create user handler")

	var request CreateUserRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode or validate request")
		return
	}

	newUserDomain := domainFromDTO(request)
	userDomain, err := h.usersService.CreateUser(ctx, newUserDomain)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to create user")
		return
	}

	userDTO := userDTOFromDomain(userDomain)
	response := CreateUserResponse(userDTO)
	responseHandler.JSONResponse(response, http.StatusCreated)
}

func domainFromDTO(dto CreateUserRequest) domain.User {
	return domain.NewUserUnitiliazied(dto.FullName, dto.PhoneNumber)
}
