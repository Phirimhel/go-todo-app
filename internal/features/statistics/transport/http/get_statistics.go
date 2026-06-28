package statistics_transport_http

import (
	"fmt"
	"net/http"
	"time"

	core_logger "github.com/Phirimhel/go-todo-app/internal/core/logger"
	core_http_response "github.com/Phirimhel/go-todo-app/internal/core/transport/http/response"
	core_http_utils "github.com/Phirimhel/go-todo-app/internal/core/transport/http/utils"
)

type GetStatistics StatisticDTOResponse

// GetStatistics godoc
// @Summary      Get statistics
// @Description  Get statistics for a specific author within a date range
// @Tags         statistics
// @Accept       json
// @Produce      json
// @Param        author_id   query     int       false  "Author ID"
// @Param        from        query     string    false  "Start date (YYYY-MM-DD)" format(date)
// @Param        to          query     string    false  "End date (YYYY-MM-DD)"   format(date)
// @Success      200         {object}  GetStatistics  "statistics successfully retrieved"
// @Failure      400         {object}  core_http_response.ErrorResponse  "Bad request"
// @Failure      404         {object}  core_http_response.ErrorResponse  "Statistics not found"
// @Failure      500         {object}  core_http_response.ErrorResponse  "Internal server error"
// @Router       /statistics [get]
func (h *StatisticsHTTPHandler) GetStatistics(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.MustGetLoggerFromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

	log.Debug("invoice get statistics handler")

	userID, from, to, err := getUserIDFromToQueryParams(r)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get user/from/to query params")
		return
	}

	statDomain, err := h.service.GetStatistics(ctx, userID, from, to)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get statistic domain")
		return
	}

	DTO := StatisticDTOfromDomain(statDomain)
	response := StatisticDTOResponse(DTO)
	responseHandler.JSONResponse(response, http.StatusCreated)
}

func getUserIDFromToQueryParams(r *http.Request) (*int, *time.Time, *time.Time, error) {

	const (
		userIDQueryParam = "author_id"
		fromQueryParam   = "from"
		toQueryParam     = "to"
	)

	userID, err := core_http_utils.GetQueryParams(r, userIDQueryParam)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("get user_id query params: %w", err)
	}

	from, err := core_http_utils.GetDateQueryParams(r, fromQueryParam)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("get 'from' query params: %w", err)
	}

	to, err := core_http_utils.GetDateQueryParams(r, toQueryParam)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("get 'to' query params: %w", err)
	}

	return userID, from, to, nil
}
