package users_transport_http

import (
	core_http_server "github.com/Phirimhel/go-todo-app/internal/core/transport/http/server"
	users_service "github.com/Phirimhel/go-todo-app/internal/features/users/service"
)

type UsersHTTPHandler struct {
	usersService users_service.UsersService
}

func (h *UsersHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method: "POST",
			Path:   "/users",
			Hanler: h.CreateUser,
		},
		{
			Method: "GET",
			Path:   "/users",
			Hanler: h.GetUsers,
		},
		{
			Method: "GET",
			Path:   "/users/{id}",
			Hanler: h.GetUser,
		},
		{
			Method: "DELETE",
			Path:   "/users/{id}",
			Hanler: h.DeleteUser,
		},
	}
}

func NewUsersHTTPHandler(service users_service.UsersService) *UsersHTTPHandler {
	return &UsersHTTPHandler{
		usersService: service,
	}
}
