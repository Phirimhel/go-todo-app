package users_transport_http

import (
	"net/http"

	core_auth "github.com/Phirimhel/go-todo-app/internal/core/auth"
	core_http_midleware "github.com/Phirimhel/go-todo-app/internal/core/transport/http/middleware"
	core_http_server "github.com/Phirimhel/go-todo-app/internal/core/transport/http/server"
	users_service "github.com/Phirimhel/go-todo-app/internal/features/users/service"
)

type UsersHTTPHandler struct {
	usersService  users_service.UsersService
	authenticator core_auth.Authenticator
}

func NewUsersHTTPHandler(service users_service.UsersService, authenticator core_auth.Authenticator) *UsersHTTPHandler {
	return &UsersHTTPHandler{
		usersService:  service,
		authenticator: authenticator,
	}
}

func (h *UsersHTTPHandler) PublicRoutes() []*core_http_server.Route {
	return []*core_http_server.Route{
		{
			Method: http.MethodPost,
			Path:   "/users/login",
			Hanler: h.LoginUser,
		},
	}
}

func (h *UsersHTTPHandler) PrivatRoutes() []*core_http_server.Route {
	routes := []*core_http_server.Route{
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

	for _, route := range routes {
		route.Middleware = append(
			route.Middleware,
			core_http_midleware.JWT(h.authenticator),
		)
	}

	return routes
}
