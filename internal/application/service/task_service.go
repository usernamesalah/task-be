package service

import (
	"context"
	"errors"
	"time"

	"task-be/internal/domain"
	"task-be/internal/infrastructure/logger"
)

type TaskServiceImpl struct {
	taskRepo domain.TaskRepository
}

func NewTaskService(taskRepo domain.TaskRepository) domain.TaskService {
	return &TaskServiceImpl{taskRepo: taskRepo}
}

func (s *TaskServiceImpl) CreateTask(ctx context.Context, title, description string) (*domain.Task, error) {
	log := logger.GetLogger()
	log.Info("Creating task", "title", title, "description", description)

	task := &domain.Task{
		Title:       title,
		Description: description,
		Status:      domain.StatusToDo,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if !task.IsValid() {
		log.Error("Invalid task data", "title", title)
		return nil, errors.New("invalid task data")
	}

	err := s.taskRepo.Create(ctx, task)
	if err != nil {
		log.Error("Failed to create task", "error", err, "title", title)
		return nil, err
	}

	log.Info("Task created successfully", "id", task.ID, "title", task.Title)
	return task, nil
}

func (s *TaskServiceImpl) GetTaskByID(ctx context.Context, id uint) (*domain.Task, error) {
	log := logger.GetLogger()
	log.Info("Getting task by ID", "id", id)

	task, err := s.taskRepo.FindByID(ctx, id)
	if err != nil {
		log.Error("Failed to get task by ID", "error", err, "id", id)
		return nil, err
	}

	log.Info("Task retrieved successfully", "id", id, "title", task.Title)
	return task, nil
}

func (s *TaskServiceImpl) GetTasks(ctx context.Context, page, limit int, status *domain.TaskStatus) ([]domain.Task, int64, error) {
	log := logger.GetLogger()
	log.Info("Getting tasks", "page", page, "limit", limit, "status", status)

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	tasks, total, err := s.taskRepo.FindAll(ctx, page, limit, status)
	if err != nil {
		log.Error("Failed to get tasks", "error", err, "page", page, "limit", limit)
		return nil, 0, err
	}

	log.Info("Tasks retrieved successfully", "count", len(tasks), "total", total)
	return tasks, total, nil
}

func (s *TaskServiceImpl) UpdateTask(ctx context.Context, id uint, title, description *string, status *domain.TaskStatus) (*domain.Task, error) {
	log := logger.GetLogger()
	log.Info("Updating task", "id", id, "title", title, "status", status)

	task, err := s.taskRepo.FindByID(ctx, id)
	if err != nil {
		log.Error("Failed to find task for update", "error", err, "id", id)
		return nil, err
	}

	if title != nil {
		task.Title = *title
	}
	if description != nil {
		task.Description = *description
	}
	if status != nil {
		if !task.IsValidStatus(*status) {
			log.Error("Invalid status provided", "status", *status, "id", id)
			return nil, errors.New("invalid status")
		}
		task.Status = *status
	}

	task.UpdatedAt = time.Now()

	if !task.IsValid() {
		log.Error("Invalid task data after update", "id", id)
		return nil, errors.New("invalid task data")
	}

	err = s.taskRepo.Update(ctx, task)
	if err != nil {
		log.Error("Failed to update task", "error", err, "id", id)
		return nil, err
	}

	log.Info("Task updated successfully", "id", id, "title", task.Title, "status", task.Status)
	return task, nil
}

func (s *TaskServiceImpl) DeleteTask(ctx context.Context, id uint) error {
	log := logger.GetLogger()
	log.Info("Deleting task", "id", id)

	err := s.taskRepo.Delete(ctx, id)
	if err != nil {
		log.Error("Failed to delete task", "error", err, "id", id)
		return err
	}

	log.Info("Task deleted successfully", "id", id)
	return nil
}
