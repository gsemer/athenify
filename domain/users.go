package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	AuthUserID uuid.UUID `gorm:"type:uuid;not null;unique" json:"auth_user_id"`
	Username   string    `gorm:"type:text;not null;unique" json:"username"`
	Role       string    `gorm:"type:text;not null" json:"role"`
	Email      string    `gorm:"type:text;not null;unique" json:"email"`
	CreatedAt  time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"updated_at"`

	Tasks []Task `gorm:"foreignKey:UserID" json:"tasks,omitempty"`
}

type UserService interface {
	Create(user User) (User, error)
	GetByID(userID uuid.UUID) (User, error)
}

type UserRepository interface {
	Create(user User) (User, error)
	GetByID(userID uuid.UUID) (User, error)
}
