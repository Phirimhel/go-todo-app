package users_transport_http

import (
	"fmt"
	"net/http"

	core_logger "github.com/Phirimhel/go-todo-app/internal/core/logger"
	core_http_response "github.com/Phirimhel/go-todo-app/internal/core/transport/http/response"
	core_http_utils "github.com/Phirimhel/go-todo-app/internal/core/transport/http/utils"
)

type GetUsersResponse []UserDTOResponse

func (h *UsersHTTPHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.GetLoggerFromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

	log.Debug("invoce get users handler")

	limit, offset, err := getLimitOffsetQueryParams(r)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get limit/offser query params")
		return
	}

	userDomains, err := h.usersService.GetUsers(ctx, limit, offset)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get users")
		return
	}

	usersDTO := usersDTOFromDomains(userDomains)
	response := GetUsersResponse(usersDTO)
	responseHandler.JSONResponse(response, http.StatusOK)
}

func getLimitOffsetQueryParams(r *http.Request) (*int, *int, error) {

	const (
		limitQueryParamsKey  = "limit"
		offsetQueryParamsKey = "offset"
	)

	limit, err := core_http_utils.GetQueryParams(r, limitQueryParamsKey)
	if err != nil {
		return nil, nil, fmt.Errorf("get limit query params: %w", err)
	}

	offser, err := core_http_utils.GetQueryParams(r, offsetQueryParamsKey)
	if err != nil {
		return nil, nil, fmt.Errorf("get offset query params: %w", err)
	}

	return limit, offser, nil
}
