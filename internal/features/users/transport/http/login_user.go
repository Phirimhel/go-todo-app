package users_transport_http

import (
	"net/http"
	"strconv"

	core_logger "github.com/Phirimhel/go-todo-app/internal/core/logger"
	core_http_request "github.com/Phirimhel/go-todo-app/internal/core/transport/http/request"
	core_http_response "github.com/Phirimhel/go-todo-app/internal/core/transport/http/response"
)

type LoginUserRequest struct {
	Email    string `json:"email" validate:"required,min=3,max=100" example:"example@gmail.com"`
	Password string `json:"password" validate:"required,min=4,max=128" example:"aboba@$_boba"`
}

type LoginUserResponse struct {
	Token string          `json:"token"`
	User  UserDTOResponse `json:"user"`
}

// LoginUser godoc
// @Summary      Login user / get token
// @Description  Get token for user by email and password
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        request body      LoginUserRequest  true  "User credentials"
// @Success      200     {object}  LoginUserResponse  "User successfully authenticated"
// @Failure      400     {object}  core_http_response.ErrorResponse  "Bad request"
// @Failure      401     {object}  core_http_response.ErrorResponse  "Unauthorized (wrong password/email)"
// @Failure      500     {object}  core_http_response.ErrorResponse  "Internal server error"
// @Router       /users/login [post]
func (h *UsersHTTPHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.MustGetLoggerFromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

	var request LoginUserRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "Failed to validate user credentials")
		return
	}

	userDomain, err := h.usersService.LoginUser(ctx, request.Email, request.Password)
	if err != nil {
		responseHandler.ErrorResponse(err, "Failed to get user")
		return
	}

	token, err := h.authenticator.MakeJWT(strconv.Itoa(userDomain.ID), userDomain.Role)
	if err != nil {
		responseHandler.ErrorResponse(err, "Failed to create JWT for user:")
		return
	}

	userDTO := userDTOFromDomain(userDomain)
	response := LoginUserResponse{
		Token: token,
		User:  UserDTOResponse(userDTO),
	}
	responseHandler.JSONResponse(response, http.StatusOK)
}
