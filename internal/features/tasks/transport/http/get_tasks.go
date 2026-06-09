package tasks_transport_http

import (
	"fmt"
	"net/http"

	core_logger "github.com/Phirimhel/go-todo-app/internal/core/logger"
	core_http_response "github.com/Phirimhel/go-todo-app/internal/core/transport/http/response"
	core_http_utils "github.com/Phirimhel/go-todo-app/internal/core/transport/http/utils"
)

type GetTasksRequest struct {
}

type GetTasksResponse []TaskDtoResponse

func (h *TasksHTTPHandler) GetTasks(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.GetLoggerFromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

	log.Debug("ivoice get tasks handle")
	authorID, limit, offset, err := getIDLimitOffsetQueryParams(r)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get limit/offser query params")
		return
	}

	tasksDomain, err := h.tasksService.GetTasks(ctx, authorID, limit, offset)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get tasks")
	}

	tasksDTO := tasksDTOfromDomain(tasksDomain)
	tasksResponse := GetTasksResponse(tasksDTO)
	responseHandler.JSONResponse(tasksResponse, http.StatusOK)

}

func getIDLimitOffsetQueryParams(r *http.Request) (*int, *int, *int, error) {

	const (
		authorIDQueryParamsKey = "author_id"
		limitQueryParamsKey    = "limit"
		offsetQueryParamsKey   = "offset"
	)

	id, err := core_http_utils.GetQueryParams(r, authorIDQueryParamsKey)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("get author_id query params: %w", err)
	}

	limit, err := core_http_utils.GetQueryParams(r, limitQueryParamsKey)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("get limit query params: %w", err)
	}

	offser, err := core_http_utils.GetQueryParams(r, offsetQueryParamsKey)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("get offset query params: %w", err)
	}

	return id, limit, offser, nil
}
