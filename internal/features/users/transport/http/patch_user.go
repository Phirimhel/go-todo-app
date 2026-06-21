package users_transport_http

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Phirimhel/go-todo-app/internal/core/domain"
	core_logger "github.com/Phirimhel/go-todo-app/internal/core/logger"
	core_http_request "github.com/Phirimhel/go-todo-app/internal/core/transport/http/request"
	core_http_response "github.com/Phirimhel/go-todo-app/internal/core/transport/http/response"
	core_http_types "github.com/Phirimhel/go-todo-app/internal/core/transport/http/types"
	core_http_utils "github.com/Phirimhel/go-todo-app/internal/core/transport/http/utils"
)

// UserPatchDocRequest is strictly used to generate clean Swagger documentation
type UserPatchDocRequest struct {
	FullName    string `json:"full_name" example:"John Doe"`
	PhoneNumber string `json:"phone_number" example:"+1234567890"`
}

type PatchUserRequest struct {
	FullName    core_http_types.Nullable[string] `json:"full_name"`
	PhoneNumber core_http_types.Nullable[string] `json:"phone_number"`
}

func (r *PatchUserRequest) Validate() error {

	// patch full_name validation
	if r.FullName.Set {
		if r.FullName.Value == nil {
			return fmt.Errorf("full_name can't be NULL")
		}

		fullNameLength := len([]rune(*r.FullName.Value))
		if fullNameLength < 3 || fullNameLength > 100 {
			return fmt.Errorf("full_name can't be less 3 or more thab 100 characters")
		}

	}

	// patch phone_number validation
	if r.PhoneNumber.Set {
		if r.PhoneNumber.Value != nil {
			phoneNumberLength := len([]rune(*r.PhoneNumber.Value))

			if phoneNumberLength < 10 || phoneNumberLength > 15 {
				return fmt.Errorf("phone_number can't be less 10 or more thab 15 characters")
			}

			if !strings.HasPrefix(*r.PhoneNumber.Value, "+") {
				return fmt.Errorf("phone_number must start with '+'")
			}

		}
	}

	return nil
}

type PatchedUserResponse UserDTOResponse

// PatchUser godoc
// @Summary      Update user
// @Description  Partially update a user by ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        id      path      string              true  "User ID to patch"
// @Param        request body      UserPatchDocRequest    true  "Fields to update"
// @Success      200     {object}  PatchedUserResponse "User successfully patched"
// @Failure      400     {object}  core_http_response.ErrorResponse "Bad request"
// @Failure      404     {object}  core_http_response.ErrorResponse "User not found"
// @Failure      500     {object}  core_http_response.ErrorResponse "Internal server error"
// @Router       /users/{id} [patch]
func (h *UsersHTTPHandler) PatchUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.GetLoggerFromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

	userID, err := core_http_utils.GetPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get user id path value")
		return
	}

	var request PatchUserRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode or validate request")
		return
	}

	userPatch := userPatchFromRequest(request)
	userDomain, err := h.usersService.PatchUser(ctx, userID, userPatch)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to patch user")
		return
	}

	userDTO := userDTOFromDomain(userDomain)
	patchedUserResponse := PatchedUserResponse(userDTO)
	responseHandler.JSONResponse(patchedUserResponse, http.StatusOK)
}

func userPatchFromRequest(request PatchUserRequest) domain.UserPatch {
	return domain.NewUserPatch(
		request.FullName.ToDomain(),
		request.PhoneNumber.ToDomain(),
	)

}
