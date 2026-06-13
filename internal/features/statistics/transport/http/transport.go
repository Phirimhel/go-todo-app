package statistics_transport_http

import (
	core_http_server "github.com/Phirimhel/go-todo-app/internal/core/transport/http/server"
	statistic_service "github.com/Phirimhel/go-todo-app/internal/features/statistics/service"
)

type StatisticsHTTPHandler struct {
	service statistic_service.StatisticService
}

func NewStatisticsHTTPHandler(service statistic_service.StatisticService) *StatisticsHTTPHandler {
	return &StatisticsHTTPHandler{
		service: service,
	}
}

func (h *StatisticsHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method: "GET",
			Path:   "/statistics",
			Hanler: h.GetStatistics,
		},
	}
}
