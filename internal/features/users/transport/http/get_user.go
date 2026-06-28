package users_transport_http

import (
	"net/http"

	core_logger "github.com/Phirimhel/go-todo-app/internal/core/logger"
	core_http_response "github.com/Phirimhel/go-todo-app/internal/core/transport/http/response"
	core_http_utils "github.com/Phirimhel/go-todo-app/internal/core/transport/http/utils"
)

type GetUserResponse UserDTOResponse

// GetUser godoc
// @Summary      Get user
// @Description  Get user by ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "User ID to get"
// @Success      204  {object}  CreateUserResponse  "User successfully geted"
// @Failure      400  {object}  core_http_response.ErrorResponse  "Bad request"
// @Failure      404  {object}  core_http_response.ErrorResponse  "User not found"
// @Failure      500  {object}  core_http_response.ErrorResponse  "Internal server error"
// @Router       /users/{id} [get]
func (h *UsersHTTPHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.MustGetLoggerFromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

	log.Debug("invoce get user handler")

	userID, err := core_http_utils.GetPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get userID path value")
		return
	}

	userDomain, err := h.usersService.GetUser(ctx, userID)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get user")
		return
	}
	userDTO := userDTOFromDomain(userDomain)
	userResponse := GetUserResponse(userDTO)
	responseHandler.JSONResponse(userResponse, http.StatusOK)
}
