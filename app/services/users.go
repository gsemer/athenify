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

// Create user
type CreateUserJob struct {
	User        domain.User
	UserService domain.UserService
	Result      chan domain.Result
}

func (job *CreateUserJob) Process() {
	user, err := job.UserService.Create(job.User)
	job.Result <- domain.Result{User: user, Error: err}
}

func (us UserService) Create(user domain.User) (domain.User, error) {
	return us.ur.Create(user)
}

// Get user by ID
type GetUserJob struct {
	UserID      uuid.UUID
	UserService domain.UserService
	Result      chan domain.Result
}

func (job *GetUserJob) Process() {
	user, err := job.UserService.GetByID(job.UserID)
	job.Result <- domain.Result{User: user, Error: err}
}

func (us UserService) GetByID(userID uuid.UUID) (domain.User, error) {
	return us.ur.GetByID(userID)
}
