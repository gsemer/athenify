package services

import (
	"athenify/domain"

	"github.com/google/uuid"
)

type UserService struct {
	ur domain.UserRepository
}

func NewUserService(ur domain.UserRepository) *UserService {
	return &UserService{ur: ur}
}

func (us UserService) Create(user domain.User) (domain.User, error) {
	return us.ur.Create(user)
}

func (us UserService) GetByID(userID uuid.UUID) (domain.User, error) {
	return us.ur.GetByID(userID)
}
