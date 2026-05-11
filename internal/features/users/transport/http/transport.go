package users_transport_http

import core_http_server "github.com/Phirimhel/go-todo-app/internal/core/transport/http/server"

type UsersHTTPHandler struct {
	usersService UsersService
}

type UsersService interface {
}

func (u *UsersHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method: "POST",
			Path:   "/users",
			Hanler: u.CreateUser,
		},
	}
}

func NewUsersHTTPHandler(service UsersService) *UsersHTTPHandler {
	return &UsersHTTPHandler{
		usersService: service,
	}
}
