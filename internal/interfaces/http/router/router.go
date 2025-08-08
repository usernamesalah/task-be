package router

import (
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	authMiddleware "task-be/internal/infrastructure/middleware"
	"task-be/internal/infrastructure/config"
	"task-be/internal/interfaces/http/handler"
)

func NewRouter(taskHandler *handler.TaskHandler, cfg *config.Config) *echo.Echo {
	e := echo.New()

	e.Use(echoMiddleware.Logger())
	e.Use(echoMiddleware.Recover())
	e.Use(echoMiddleware.CORS())

	tasks := e.Group("/tasks")
	tasks.GET("", taskHandler.GetTasks)
	tasks.GET("/:id", taskHandler.GetTaskByID)

	authTasks := e.Group("/tasks")
	authTasks.Use(authMiddleware.BasicAuth(cfg))
	authTasks.POST("", taskHandler.CreateTask)
	authTasks.PATCH("/:id", taskHandler.UpdateTask)
	authTasks.DELETE("/:id", taskHandler.DeleteTask)

	return e
}
