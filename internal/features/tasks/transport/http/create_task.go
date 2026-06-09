package tasks_transport_http

import (
	"net/http"

	"github.com/Phirimhel/go-todo-app/internal/core/domain"
	core_logger "github.com/Phirimhel/go-todo-app/internal/core/logger"
	core_http_request "github.com/Phirimhel/go-todo-app/internal/core/transport/http/request"
	core_http_response "github.com/Phirimhel/go-todo-app/internal/core/transport/http/response"
)

type CreateTaskRequest struct {
	Title       string  `json:"title" validate:"required,min=1,max=100"`
	Description *string `json:"description" validate:"omitempty,min=1,max=1000"`
	Completed   bool    `json:"completed" default:"false"`
	AuthorID    int     `json:"author_id" validate:"required"`
}

type createTaskResponse TaskDtoResponse

func (h *TasksHTTPHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.GetLoggerFromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

	log.Debug("invoce create task handler")

	var request CreateTaskRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode/validate json")
		return
	}

	newTaskDomain := domainFromDTO(request)
	taksDomain, err := h.tasksService.CreateTask(ctx, newTaskDomain)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to create task")
		return
	}

	taskDTO := taskDTOfromDomain(taksDomain)
	taskResponse := createTaskResponse(taskDTO)
	responseHandler.JSONResponse(taskResponse, http.StatusCreated)

}

func domainFromDTO(dto CreateTaskRequest) domain.Task {
	return domain.NewTaskUnitialized(
		dto.Title,
		dto.Description,
		dto.Completed,
		dto.AuthorID,
	)
}
