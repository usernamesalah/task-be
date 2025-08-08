package dto

type CreateTaskRequest struct {
	Title       string `json:"title" validate:"required,max=255"`
	Description string `json:"description"`
}

type UpdateTaskRequest struct {
	Title       *string `json:"title" validate:"omitempty,max=255"`
	Description *string `json:"description"`
	Status      *string `json:"status" validate:"omitempty,oneof=TO_DO IN_PROGRESS DONE"`
}

type TaskResponse struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type TaskListResponse struct {
	Tasks []TaskResponse `json:"tasks"`
	Total int64         `json:"total"`
	Page  int           `json:"page"`
	Limit int           `json:"limit"`
}
