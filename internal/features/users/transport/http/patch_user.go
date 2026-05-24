package users_transport_http

import (
	"fmt"
	"net/http"

	core_logger "github.com/Phirimhel/go-todo-app/internal/core/logger"
	core_http_request "github.com/Phirimhel/go-todo-app/internal/core/transport/http/request"
	core_http_response "github.com/Phirimhel/go-todo-app/internal/core/transport/http/response"
	core_http_types "github.com/Phirimhel/go-todo-app/internal/core/transport/http/types"
	core_http_utils "github.com/Phirimhel/go-todo-app/internal/core/transport/http/utils"
)

type PatchUserRequest struct {
	FullName    core_http_types.Nullable[string] `json:"full_name"`
	PhoneNumber core_http_types.Nullable[string] `json:"phone_number"`
}

type PatchUserResponse UserDTOResponse

func (h *UsersHTTPHandler) PatchUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.GetLoggerFromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

	userID, err := core_http_utils.GetPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get user id path value")
	}

	var request PatchUserRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode or validate request")
		return
	}

	fmt.Println("request ****************", request)
	w.WriteHeader(http.StatusOK)

	//log.Debug("invoce patch user handler")

}
