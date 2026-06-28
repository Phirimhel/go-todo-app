package users_transport_http

import (
	"net/http"

	core_logger "github.com/Phirimhel/go-todo-app/internal/core/logger"
	core_http_request "github.com/Phirimhel/go-todo-app/internal/core/transport/http/request"
	core_http_response "github.com/Phirimhel/go-todo-app/internal/core/transport/http/response"
)

type LoginUserRequest struct {
	FullName string `json:"full_name"`
	Password string `json:"password"`
}

type LoginUserResponse struct {
	Token string `json:"token"`
}

func (h *UsersHTTPHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.GetLoggerFromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

	var loginData LoginUserRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &loginData); err != nil {
		responseHandler.ErrorResponse(err, "Failed to validate user credentials")
	}

	token, err := h.usersService.LoginUser(ctx, loginData.FullName, loginData.Password)
	if err != nil {
		responseHandler.ErrorResponse(err, "Failed to get token for user")
	}

	response := LoginUserResponse{
		Token: token,
	}
	responseHandler.JSONResponse(response, http.StatusCreated)
}
