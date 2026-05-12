package users_transport_http

import (
	"encoding/json"
	"net/http"

	core_logger "github.com/Phirimhel/go-todo-app/internal/core/logger"
)

type CreateUserRequest struct {
	FullName    string  `json:"full_name"`
	PhoneNumber *string `json:"phone_number"`
}

type CreateUserResponse struct {
	ID          int     `json:"id"`
	Version     int     `json:"version"`
	FullName    string  `json:"full_name"`
	PhoneNumber *string `json:"phone_number"`
}

func (u *UsersHTTPHandler) CreateUser(w http.ResponseWriter, r *http.Request) {

	log := core_logger.FromContext(r.Context())

	log.Debug("invoce create user handler")

	var request CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return
	}

	w.WriteHeader(http.StatusAccepted)

}
