package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"task-be/internal/application/service"
	"task-be/internal/infrastructure/config"
	"task-be/internal/infrastructure/database"
	"task-be/internal/infrastructure/logger"
	"task-be/internal/infrastructure/repository"
	"task-be/internal/interfaces/http/handler"
	"task-be/internal/interfaces/http/router"

	"github.com/joho/godotenv"
)

func main() {
	// Initialize logger
	logger.InitLogger()
	log := logger.GetLogger()

	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Info("No .env file found, using default environment variables")
	}

	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Error("Failed to load configuration", "error", err)
		panic("Failed to load configuration")
	}

	// Create context with cancellation
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Initialize database
	db := database.NewDatabase(ctx, cfg)
	taskRepo := repository.NewTaskRepository(db)
	taskService := service.NewTaskService(taskRepo)
	taskHandler := handler.NewTaskHandler(taskService)

	// Initialize router
	e := router.NewRouter(taskHandler, cfg)

	// Start server in a goroutine
	go func() {
		log.Info("Server starting", "port", cfg.Server.Port)
		if err := e.Start(":" + cfg.Server.Port); err != nil {
			log.Error("Failed to start server", "error", err)
			cancel()
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Info("Shutting down server...")

	// Create shutdown context with timeout
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer shutdownCancel()

	// Gracefully shutdown the server
	if err := e.Shutdown(shutdownCtx); err != nil {
		log.Error("Server forced to shutdown", "error", err)
	}

	log.Info("Server exited")
}
