package persistence

import (
	"athenify/domain"
	"log"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ur UserRepository) Create(user domain.User) (domain.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Unable to hash password: %s", err.Error())
		return domain.User{}, err
	}
	user.Password = string(hashedPassword)
	result := ur.db.Create(&user)
	if result.Error != nil {
		log.Printf("Unable to create user: %s", result.Error)
		return domain.User{}, result.Error
	}
	return user, nil
}

func (ur UserRepository) GetByID(userID uuid.UUID) (domain.User, error) {
	var user domain.User
	result := ur.db.First(&user, "id = ?", userID)
	if result.Error != nil {
		log.Printf("User not found: %s", result.Error)
		return domain.User{}, result.Error
	}
	return user, nil
}
