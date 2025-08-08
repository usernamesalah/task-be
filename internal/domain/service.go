package domain

import "context"

type TaskService interface {
	CreateTask(ctx context.Context, title, description string) (*Task, error)
	GetTaskByID(ctx context.Context, id uint) (*Task, error)
	GetTasks(ctx context.Context, page, limit int, status *TaskStatus) ([]Task, int64, error)
	UpdateTask(ctx context.Context, id uint, title, description *string, status *TaskStatus) (*Task, error)
	DeleteTask(ctx context.Context, id uint) error
}
