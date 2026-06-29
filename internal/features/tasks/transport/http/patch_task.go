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
	Title       core_http_types.Nullable[string] `json:"title" swaggertype:"string" example:"Deploy auth service to staging"`
	Description core_http_types.Nullable[string] `json:"description" swaggertype:"string" example:"Build the docker image, run migrations, and update the Kubernetes deployment configuration."`
	Completed   core_http_types.Nullable[bool]   `json:"completed" swaggertype:"boolean" example:"true"`
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

// PatchTask godoc
// @Summary      Update task
// @Description  Partially update a task by ID
// @Description  ### Update fields logic (three-state logic):
// @Description  **Field is not passed**: `description` or `title` is ignored; the value in the DB will not change.
// @Description  **Field is passed**: `"description": "Read a book"`, the value in the DB will change.
// @Description  **Field is passed as null**: `"description": null`, the value in the DB will be deleted (or set to null).
// @Description  **Constraints**: The `title` field cannot be set to null.
// @Tags         tasks
// @Accept       json
// @Produce      json
// @Param        id      path      string              true  "Task ID to patch"
// @Param        request body      PatchTaskRequest    true  "Fields to update"
// @Success      200     {object}  PatchTaskResponse   "Task successfully patched"
// @Failure      400     {object}  core_http_response.ErrorResponse "Bad request"
// @Failure      404     {object}  core_http_response.ErrorResponse "Task not found"
// @Failure      500     {object}  core_http_response.ErrorResponse "Internal server error"
// @Router       /tasks/{id} [patch]
func (h *TasksHTTPHandler) PatchTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.MustGetLoggerFromContext(ctx)
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
