package handler

import (
	"net/http"
	"strconv"
	"time"

	"task-be/internal/domain"
	"task-be/internal/interfaces/http/dto"

	"github.com/labstack/echo/v4"
)

type TaskHandler struct {
	taskService domain.TaskService
}

func NewTaskHandler(taskService domain.TaskService) *TaskHandler {
	return &TaskHandler{taskService: taskService}
}

func (h *TaskHandler) CreateTask(c echo.Context) error {
	var req dto.CreateTaskRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
	}

	if err := c.Validate(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	task, err := h.taskService.CreateTask(c.Request().Context(), req.Title, req.Description)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	response := dto.TaskResponse{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Status:      string(task.Status),
		CreatedAt:   task.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   task.UpdatedAt.Format(time.RFC3339),
	}

	return c.JSON(http.StatusCreated, response)
}

func (h *TaskHandler) GetTasks(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	status := c.QueryParam("status")

	var taskStatus *domain.TaskStatus
	if status != "" {
		ts := domain.TaskStatus(status)
		taskStatus = &ts
	}

	tasks, total, err := h.taskService.GetTasks(c.Request().Context(), page, limit, taskStatus)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	var taskResponses []dto.TaskResponse
	for _, task := range tasks {
		taskResponses = append(taskResponses, dto.TaskResponse{
			ID:          task.ID,
			Title:       task.Title,
			Description: task.Description,
			Status:      string(task.Status),
			CreatedAt:   task.CreatedAt.Format(time.RFC3339),
			UpdatedAt:   task.UpdatedAt.Format(time.RFC3339),
		})
	}

	response := dto.TaskListResponse{
		Tasks: taskResponses,
		Total: total,
		Page:  page,
		Limit: limit,
	}

	return c.JSON(http.StatusOK, response)
}

func (h *TaskHandler) GetTaskByID(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid task ID")
	}

	task, err := h.taskService.GetTaskByID(c.Request().Context(), uint(id))
	if err != nil {
		if err.Error() == "task not found" {
			return echo.NewHTTPError(http.StatusNotFound, "Task not found")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	response := dto.TaskResponse{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Status:      string(task.Status),
		CreatedAt:   task.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   task.UpdatedAt.Format(time.RFC3339),
	}

	return c.JSON(http.StatusOK, response)
}

func (h *TaskHandler) UpdateTask(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid task ID")
	}

	var req dto.UpdateTaskRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body")
	}

	if err := c.Validate(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var taskStatus *domain.TaskStatus
	if req.Status != nil {
		ts := domain.TaskStatus(*req.Status)
		taskStatus = &ts
	}

	task, err := h.taskService.UpdateTask(c.Request().Context(), uint(id), req.Title, req.Description, taskStatus)
	if err != nil {
		if err.Error() == "task not found" {
			return echo.NewHTTPError(http.StatusNotFound, "Task not found")
		}
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	response := dto.TaskResponse{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Status:      string(task.Status),
		CreatedAt:   task.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   task.UpdatedAt.Format(time.RFC3339),
	}

	return c.JSON(http.StatusOK, response)
}

func (h *TaskHandler) DeleteTask(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid task ID")
	}

	err = h.taskService.DeleteTask(c.Request().Context(), uint(id))
	if err != nil {
		if err.Error() == "task not found" {
			return echo.NewHTTPError(http.StatusNotFound, "Task not found")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}
