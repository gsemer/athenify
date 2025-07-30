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
	user, err := us.ur.Create(user)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (us UserService) GetByID(userID uuid.UUID) (domain.User, error) {
	user, err := us.ur.GetByID(userID)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}
