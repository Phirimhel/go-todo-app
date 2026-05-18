package users_transport_http

import (
	core_http_server "github.com/Phirimhel/go-todo-app/internal/core/transport/http/server"
	users_service "github.com/Phirimhel/go-todo-app/internal/features/users/service"
)

type UsersHTTPHandler struct {
	usersService users_service.UsersService
}

func (u *UsersHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method: "POST",
			Path:   "/users",
			Hanler: u.CreateUser,
		},
		{
			Method: "GET",
			Path:   "/users",
			Hanler: u.CreateUser,
		},
	}
}

func NewUsersHTTPHandler(service users_service.UsersService) *UsersHTTPHandler {
	return &UsersHTTPHandler{
		usersService: service,
	}
}
