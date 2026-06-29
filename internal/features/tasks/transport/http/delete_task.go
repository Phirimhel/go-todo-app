package tasks_transport_http

import (
	"net/http"

	core_logger "github.com/Phirimhel/go-todo-app/internal/core/logger"
	core_http_response "github.com/Phirimhel/go-todo-app/internal/core/transport/http/response"
	core_http_utils "github.com/Phirimhel/go-todo-app/internal/core/transport/http/utils"
)

// DeleteTask godoc
// @Summary      Delete task
// @Description  Delete task by ID
// @Tags         tasks
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "task ID to delete"
// @Success      204  "task successfully deleted"
// @Failure      400  {object}  core_http_response.ErrorResponse  "Bad request"
// @Failure      404  {object}  core_http_response.ErrorResponse  "task not found"
// @Failure      500  {object}  core_http_response.ErrorResponse  "Internal server error"
// @Router       /tasks/{id} [delete]
func (h TasksHTTPHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.MustGetLoggerFromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

	log.Debug("invoice delete task handler")

	taskID, err := core_http_utils.GetPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "${failed to get ID from path value}")
		return
	}

	if err := h.tasksService.DeleteTask(ctx, taskID); err != nil {
		responseHandler.ErrorResponse(err, "failed to delete task")
		return
	}

	responseHandler.NoContentResponse(http.StatusNoContent)
}
