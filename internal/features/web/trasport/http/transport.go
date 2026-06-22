package web_transport_http

import (
	core_http_server "github.com/Phirimhel/go-todo-app/internal/core/transport/http/server"
	web_service "github.com/Phirimhel/go-todo-app/internal/features/web/service"
	web_service_http "github.com/Phirimhel/go-todo-app/internal/features/web/service"
)

type WebHTTPHandler struct {
	webService web_service_http.WebService
}

func NewWebHTTPHandler(service web_service.WebService) *WebHTTPHandler {
	return &WebHTTPHandler{
		webService: service,
	}
}

func (h *WebHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Path:   "/",
			Hanler: h.GetMainPage,
		},
	}
}
