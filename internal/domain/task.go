package domain

import (
	"time"
)

type TaskStatus string

const (
	StatusToDo       TaskStatus = "TO_DO"
	StatusInProgress TaskStatus = "IN_PROGRESS"
	StatusDone       TaskStatus = "DONE"
)

type Task struct {
	ID          uint       `json:"id" gorm:"primaryKey"`
	Title       string     `json:"title" gorm:"size:255;not null"`
	Description string     `json:"description" gorm:"type:text"`
	Status      TaskStatus `json:"status" gorm:"default:'TO_DO'"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

func (t *Task) IsValid() bool {
	return len(t.Title) > 0 && len(t.Title) <= 255
}

func (t *Task) IsValidStatus(status TaskStatus) bool {
	return status == StatusToDo || status == StatusInProgress || status == StatusDone
}
