package core_http_server

import "net/http"

type Route struct {
	Method string
	Path   string
	Hanler http.HandlerFunc
}

func NewRoute(Method, Path string, Hanler http.HandlerFunc) Route {
	return Route{
		Method: Method,
		Path:   Path,
		Hanler: Hanler,
	}
}
