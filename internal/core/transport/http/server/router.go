package core_http_server

import (
	"fmt"
	"net/http"

	core_http_midleware "github.com/Phirimhel/go-todo-app/internal/core/transport/http/middleware"
)

type ApiVersion string

var (
	ApiVersion1 = ApiVersion("v1")
	ApiVersion2 = ApiVersion("v2")
	ApiVersion3 = ApiVersion("v3")
)

type ApiVersionRouter struct {
	*http.ServeMux
	ApiVersion ApiVersion
	Middleware []core_http_midleware.Middleware
}

func NewApiVersionRouter(version ApiVersion, middleware ...core_http_midleware.Middleware) *ApiVersionRouter {
	return &ApiVersionRouter{
		ServeMux:   &http.ServeMux{},
		ApiVersion: version,
		Middleware: middleware,
	}
}

func (r *ApiVersionRouter) RegisterRoutes(routes ...Route) {
	for _, route := range routes {
		pattern := fmt.Sprintf("%s %s", route.Method, route.Path)

		fmt.Println(pattern)
		r.ServeMux.Handle(pattern, route.WithMiddleware())
	}
}

func (r *ApiVersionRouter) WithMiddleware() http.Handler {
	return core_http_midleware.ChaneMiddleware(
		r.ServeMux,
		r.Middleware...,
	)
}
