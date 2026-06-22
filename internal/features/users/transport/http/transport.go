package users_transport_http

import (
	"net/http"

	core_http_midleware "github.com/Phirimhel/go-todo-app/internal/core/transport/http/middleware"
	core_http_server "github.com/Phirimhel/go-todo-app/internal/core/transport/http/server"
	users_service "github.com/Phirimhel/go-todo-app/internal/features/users/service"
)

type UsersHTTPHandler struct {
	usersService users_service.UsersService
}

func NewUsersHTTPHandler(service users_service.UsersService) *UsersHTTPHandler {
	return &UsersHTTPHandler{
		usersService: service,
	}
}

func (h *UsersHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method: http.MethodPost,
			Path:   "/users",
			Hanler: h.CreateUser,
		},
		{
			Method: http.MethodGet,
			Path:   "/users",
			Hanler: h.GetUsers,
		},
		{
			Method: http.MethodGet,
			Path:   "/users/{id}",
			Hanler: h.GetUser,
			Middleware: []core_http_midleware.Middleware{
				core_http_midleware.MockMiddleware(),
			},
		},
		{
			Method: http.MethodDelete,
			Path:   "/users/{id}",
			Hanler: h.DeleteUser,
		},
		{
			Method: http.MethodPatch,
			Path:   "/users/{id}",
			Hanler: h.PatchUser,
		},
	}
}
