package users_transport_http

import (
	"net/http"
	"strconv"

	core_logger "github.com/Phirimhel/go-todo-app/internal/core/logger"
	core_http_response "github.com/Phirimhel/go-todo-app/internal/core/transport/http/response"
)

type GetUserResponse UserDTOResponse

func (h *UsersHTTPHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)
	userIDString := r.PathValue("id")

	log.Debug("invoce get user handler")

	userID, err := strconv.Atoi(userIDString)
	if err != nil {
		responseHandler.ErrorResponse(err, "[handler]: failed to parse or validate user id")
	}

	userDomain, err := h.usersService.GetUser(ctx, userID)
	if err != nil {
		responseHandler.ErrorResponse(err, "[handler]: failed to get user domain from service")
	}

	userDTO := userDTOFromDomain(userDomain)
	userResponse := GetUserResponse(userDTO)
	responseHandler.JSONResponse(userResponse, http.StatusOK)
}
