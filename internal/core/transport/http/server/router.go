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
}

func NewApiVersionRouter(version ApiVersion) *ApiVersionRouter {
	return &ApiVersionRouter{
		ServeMux:   &http.ServeMux{},
		ApiVersion: version,
	}
}

func (a *ApiVersionRouter) RegisterRoutes(routes ...Route) {
	for _, route := range routes {
		pattern := fmt.Sprintf("%s %s", route.Method, route.Path)

		h := core_http_midleware.ChaneMiddleware(route.Hanler, route.Middleware...)
		a.ServeMux.Handle(pattern, h)
	}
}
