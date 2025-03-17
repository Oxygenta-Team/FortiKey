package service

import (
	"github.com/Oxygenta-Team/FortiKey/pkg/user-management/repository"

	"github.com/Oxygenta-Team/FortiKey/pkg/models"
)

type UserService struct {
	repo *repository.Repository
}

func NewUserService(repo *repository.Repository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(login, email string) (*models.User, error) {
	return s.repo.CreateUser(login, email)
}

func (s *UserService) DeleteUser(id int) error {
	return s.repo.DeleteUser(id)
}

func (s *UserService) GetUserByID(id int) (*models.User, error) {
	return s.repo.GetUserByID(id)
}
