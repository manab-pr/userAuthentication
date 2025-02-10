package services

import (
	"user-authentication/models"
	"user-authentication/repositories"
)

type UserService struct {
	UserRepo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{UserRepo: repo}

}

func (s *UserService) Getuser(id uint) (*models.User, error) {
	return s.UserRepo.GetUserByID(id)

}

func (s *UserService) CreateUser(name string, email string) (*models.User, error) {
	user := &models.User{Name: name, Email: email}
	err := s.UserRepo.CreateUser(user)

	return user, err
}

func (s *UserService) ListUsers() ([]models.User, error) {
	return s.UserRepo.ListUsers()
}
