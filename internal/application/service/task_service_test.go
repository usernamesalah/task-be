package service

import (
	"context"
	"errors"
	"testing"
	"time"

	"task-be/internal/domain"
)

type MockTaskRepository struct {
	tasks  map[uint]*domain.Task
	nextID uint
}

func NewMockTaskRepository() *MockTaskRepository {
	return &MockTaskRepository{
		tasks:  make(map[uint]*domain.Task),
		nextID: 1,
	}
}

func (m *MockTaskRepository) Create(ctx context.Context, task *domain.Task) error {
	task.ID = m.nextID
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()
	m.tasks[task.ID] = task
	m.nextID++
	return nil
}

func (m *MockTaskRepository) FindByID(ctx context.Context, id uint) (*domain.Task, error) {
	task, exists := m.tasks[id]
	if !exists {
		return nil, errors.New("task not found")
	}
	return task, nil
}

func (m *MockTaskRepository) FindAll(ctx context.Context, page, limit int, status *domain.TaskStatus) ([]domain.Task, int64, error) {
	var tasks []domain.Task
	for _, task := range m.tasks {
		if status == nil || task.Status == *status {
			tasks = append(tasks, *task)
		}
	}
	return tasks, int64(len(tasks)), nil
}

func (m *MockTaskRepository) Update(ctx context.Context, task *domain.Task) error {
	if _, exists := m.tasks[task.ID]; !exists {
		return errors.New("task not found")
	}
	task.UpdatedAt = time.Now()
	m.tasks[task.ID] = task
	return nil
}

func (m *MockTaskRepository) Delete(ctx context.Context, id uint) error {
	if _, exists := m.tasks[id]; !exists {
		return errors.New("task not found")
	}
	delete(m.tasks, id)
	return nil
}

func TestCreateTask(t *testing.T) {
	repo := NewMockTaskRepository()
	service := NewTaskService(repo)
	ctx := context.Background()

	task, err := service.CreateTask(ctx, "Test Task", "Test Description")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if task.Title != "Test Task" {
		t.Errorf("Expected title 'Test Task', got %s", task.Title)
	}

	if task.Status != domain.StatusToDo {
		t.Errorf("Expected status 'TO_DO', got %s", task.Status)
	}
}

func TestCreateTaskWithEmptyTitle(t *testing.T) {
	repo := NewMockTaskRepository()
	service := NewTaskService(repo)
	ctx := context.Background()

	_, err := service.CreateTask(ctx, "", "Test Description")
	if err == nil {
		t.Error("Expected error for empty title")
	}
}

func TestGetTaskByID(t *testing.T) {
	repo := NewMockTaskRepository()
	service := NewTaskService(repo)
	ctx := context.Background()

	createdTask, _ := service.CreateTask(ctx, "Test Task", "Test Description")
	retrievedTask, err := service.GetTaskByID(ctx, createdTask.ID)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if retrievedTask.ID != createdTask.ID {
		t.Errorf("Expected ID %d, got %d", createdTask.ID, retrievedTask.ID)
	}
}

func TestGetTaskByIDNotFound(t *testing.T) {
	repo := NewMockTaskRepository()
	service := NewTaskService(repo)
	ctx := context.Background()

	_, err := service.GetTaskByID(ctx, 999)
	if err == nil {
		t.Error("Expected error for non-existent task")
	}
}
