package repository

import (
	"context"
	"errors"

	"task-be/internal/domain"
	"task-be/internal/infrastructure/logger"

	"gorm.io/gorm"
)

type TaskRepositoryImpl struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) domain.TaskRepository {
	return &TaskRepositoryImpl{db: db}
}

func (r *TaskRepositoryImpl) Create(ctx context.Context, task *domain.Task) error {
	log := logger.GetLogger()
	log.Info("Creating task", "title", task.Title, "status", task.Status)
	return r.db.WithContext(ctx).Create(task).Error
}

func (r *TaskRepositoryImpl) FindByID(ctx context.Context, id uint) (*domain.Task, error) {
	var task domain.Task
	err := r.db.WithContext(ctx).First(&task, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("task not found")
		}
		return nil, err
	}
	return &task, nil
}

func (r *TaskRepositoryImpl) FindAll(ctx context.Context, page, limit int, status *domain.TaskStatus) ([]domain.Task, int64, error) {
	var tasks []domain.Task
	var total int64

	query := r.db.WithContext(ctx).Model(&domain.Task{})
	if status != nil {
		query = query.Where("status = ?", *status)
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	err = query.Offset(offset).Limit(limit).Find(&tasks).Error
	if err != nil {
		return nil, 0, err
	}

	return tasks, total, nil
}

func (r *TaskRepositoryImpl) Update(ctx context.Context, task *domain.Task) error {
	log := logger.GetLogger()
	log.Info("Updating task", "id", task.ID, "title", task.Title, "status", task.Status)
	return r.db.WithContext(ctx).Save(task).Error
}

func (r *TaskRepositoryImpl) Delete(ctx context.Context, id uint) error {
	log := logger.GetLogger()
	log.Info("Deleting task", "id", id)
	result := r.db.WithContext(ctx).Delete(&domain.Task{}, id)
	if result.RowsAffected == 0 {
		return errors.New("task not found")
	}
	return result.Error
}
