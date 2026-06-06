package core_http_server

import (
	"net/http"

	core_http_midleware "github.com/Phirimhel/go-todo-app/internal/core/transport/http/middleware"
)

type Route struct {
	Method     string
	Path       string
	Hanler     http.HandlerFunc
	Middleware []core_http_midleware.Middleware
}

func NewRoute(Method, Path string, Hanler http.HandlerFunc) Route {
	return Route{
		Method: Method,
		Path:   Path,
		Hanler: Hanler,
	}
}
