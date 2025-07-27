package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Username  string    `gorm:"type:text;not null;unique" json:"username"`
	Password  string    `gorm:"type:text;not null" json:"password"`
	Role      string    `gorm:"type:text;not null" json:"role"`
	Email     string    `gorm:"type:text;not null;unique" json:"email"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	Tasks []Task `gorm:"foreignKey:UserID" json:"tasks,omitempty"`
}
