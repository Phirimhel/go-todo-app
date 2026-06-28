package users_transport_http

import (
	"net/http"

	core_auth "github.com/Phirimhel/go-todo-app/internal/core/auth"
	core_logger "github.com/Phirimhel/go-todo-app/internal/core/logger"
	core_http_response "github.com/Phirimhel/go-todo-app/internal/core/transport/http/response"
)

// type LoginUserRequest struct {
// 	FullName string `json:"full_name" validate:"required,min=3,max=100" example:"John Doe"`
// 	Password string `json:"password" validate:"required,min=4,max=128" example:"aboba@$_boba"`
// }

type LoginUserResponse struct {
	Token string `json:"token"`
}

func (h *UsersHTTPHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.MustGetLoggerFromContext(ctx)

	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

	// var loginData LoginUserRequest
	// if err := core_http_request.DecodeAndValidateRequest(r, &loginData); err != nil {
	// 	responseHandler.ErrorResponse(err, "Failed to validate user credentials")
	// }

	// token, err := h.usersService.LoginUser(ctx, loginData.FullName, loginData.Password)
	// if err != nil {
	// 	responseHandler.ErrorResponse(err, "Failed to get token for user")
	// }

	authService := core_auth.NewAuthenticator(core_auth.NewConfigMust())
	token, _ := authService.MakeJWT("1", "user")

	response := LoginUserResponse{
		Token: token,
	}
	responseHandler.JSONResponse(response, http.StatusOK)
}
