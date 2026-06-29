package web_transport_http

import (
	"net/http"

	core_logger "github.com/Phirimhel/go-todo-app/internal/core/logger"
	core_http_response "github.com/Phirimhel/go-todo-app/internal/core/transport/http/response"
)

func (h *WebHTTPHandler) GetMainPage(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.MustGetLoggerFromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)
	log.Debug("invoice get main page index.http")

	mainPage, err := h.webService.GetMainPage()
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get index.html for main page")
	}

	responseHandler.HTMLResponse(mainPage)
}
