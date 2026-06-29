package tasks_transport_http

import (
	core_http_server "github.com/Phirimhel/go-todo-app/internal/core/transport/http/server"
	tasks_service "github.com/Phirimhel/go-todo-app/internal/features/tasks/service"
)

type TasksHTTPHandler struct {
	tasksService tasks_service.TasksService
}

func NewTasksHTTPHandler(service tasks_service.TasksService) *TasksHTTPHandler {
	return &TasksHTTPHandler{
		tasksService: service,
	}
}

func (h *TasksHTTPHandler) PrivetRoutes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method: "POST",
			Path:   "/tasks",
			Hanler: h.CreateTask,
		},
		{
			Method: "GET",
			Path:   "/tasks",
			Hanler: h.GetTasks,
		},
		{
			Method: "GET",
			Path:   "/tasks/{id}",
			Hanler: h.GetTask,
		},
		{
			Method: "DELETE",
			Path:   "/tasks/{id}",
			Hanler: h.DeleteTask,
		},
		{
			Method: "PATCH",
			Path:   "/tasks/{id}",
			Hanler: h.PatchTask,
		},
	}
}
