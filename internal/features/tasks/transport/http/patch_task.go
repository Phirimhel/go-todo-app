package tasks_transport_http

import (
	"fmt"
	"net/http"

	"github.com/Phirimhel/go-todo-app/internal/core/domain"
	core_logger "github.com/Phirimhel/go-todo-app/internal/core/logger"
	core_http_request "github.com/Phirimhel/go-todo-app/internal/core/transport/http/request"
	core_http_response "github.com/Phirimhel/go-todo-app/internal/core/transport/http/response"
	core_http_types "github.com/Phirimhel/go-todo-app/internal/core/transport/http/types"
	core_http_utils "github.com/Phirimhel/go-todo-app/internal/core/transport/http/utils"
)

type PatchTaskRequest struct {
	Title       core_http_types.Nullable[string] `json:"title"`
	Description core_http_types.Nullable[string] `json:"description"`
	Completed   core_http_types.Nullable[bool]   `json:"completed"`
}

func (r *PatchTaskRequest) Validate() error {

	// title validation
	if r.Title.Set {
		if r.Title.Value == nil {
			return fmt.Errorf("title can't be NULL")
		}
		titleLength := len([]rune(*r.Title.Value))
		if titleLength < 3 || titleLength > 100 {
			return fmt.Errorf("title of task can't be less 3 or more thab 100 characters")
		}
	}

	// description validation
	if r.Description.Set {
		if r.Description.Value != nil {
			descriptionLength := len([]rune(*r.Description.Value))
			if descriptionLength < 1 || descriptionLength > 1000 {
				return fmt.Errorf("description of task can't be less 1 or more than 1000 characters")
			}
		}
	}

	// completed validation
	if r.Completed.Set {
		if r.Completed.Value == nil {
			return fmt.Errorf("completed can't be NULL, true/false only")
		}

	}

	return nil
}

type PatchTaskResponse TaskDtoResponse

func (h *TasksHTTPHandler) PatchTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.GetLoggerFromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

	log.Debug("invoice patch task handler")

	taskID, err := core_http_utils.GetPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get ID from path value")
		return
	}

	var request PatchTaskRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode or validate request")
		return
	}

	taskPatch := taskPatchFromRequest(request)
	taskDomain, err := h.tasksService.PatchTask(ctx, taskID, taskPatch)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get task domen")
		return
	}

	taskDTO := taskDTOfromDomain(taskDomain)
	Response := PatchTaskResponse(taskDTO)
	responseHandler.JSONResponse(Response, http.StatusCreated)
}

func taskPatchFromRequest(request PatchTaskRequest) domain.TaskPatch {
	return domain.TaskPatch{
		Title:       request.Title.ToDomain(),
		Description: request.Description.ToDomain(),
		Completed:   request.Completed.ToDomain(),
	}
}
