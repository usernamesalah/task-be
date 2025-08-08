package domain

import "context"

type TaskRepository interface {
	Create(ctx context.Context, task *Task) error
	FindByID(ctx context.Context, id uint) (*Task, error)
	FindAll(ctx context.Context, page, limit int, status *TaskStatus) ([]Task, int64, error)
	Update(ctx context.Context, task *Task) error
	Delete(ctx context.Context, id uint) error
}
