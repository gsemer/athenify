package domain

import (
	"time"

	"github.com/google/uuid"
)

type TaskStatus string

const (
	StatusPending    TaskStatus = "pending"
	StatusInProgress TaskStatus = "in-progress"
	StatusCompleted  TaskStatus = "completed"
)

type TaskPriority string

const (
	PriorityLow    TaskPriority = "low"
	PriorityMedium TaskPriority = "medium"
	PriorityHigh   TaskPriority = "high"
)

type Task struct {
	ID          uuid.UUID    `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	UserID      uuid.UUID    `gorm:"type:uuid;not null" json:"user_id"`
	Title       string       `gorm:"type:text;unique;not null" json:"title"`
	Description string       `gorm:"type:text" json:"description"`
	Status      TaskStatus   `gorm:"type:text;default:'pending';not null" json:"status"`
	Priority    TaskPriority `gorm:"type:text;default:'medium';not null" json:"priority"`
	CreatedAt   time.Time    `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time    `gorm:"autoUpdateTime" json:"updated_at"`

	User *User `gorm:"foreignKey:UserID;references:ID" json:"user"`
}
