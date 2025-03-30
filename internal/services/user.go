package services

import (
	"fmt"
	"time"

	"github.com/CaioProg/GoFinance/internal/models"
	"github.com/CaioProg/GoFinance/internal/repositories"
)

type UserService struct {
	UserRepository repositories.UserRepository
}

func (s *UserService) GetUserById(id uint) (*models.User, error) {
	return s.UserRepository.GetUserById(id)
}

func (s *UserService) GetAllUsers() (*[]models.User, error) {
	return s.UserRepository.GetAllUsers()
}

func (s *UserService) CreateUser(user *models.User) (*models.User, error) {

	existEmail, err := s.UserRepository.GetUserByEmail(user.Email)
	if err != nil {
		return nil, err
	}

	if existEmail != nil {
		return nil, fmt.Errorf("email '%s' already exists", user.Email)
	}

	return s.UserRepository.CreateUser(user)
}

func (s *UserService) UpdateUser(user *models.User, id uint) (*models.User, error) {
	var updateUser models.User
	updateUser.UpdatedAt = time.Now()

	if user.Name != "" {
		updateUser.Name = user.Name
	}

	if user.Password != "" {
		updateUser.Password = user.Password
	}

	updatedUser, err := s.UserRepository.UpdateUser(&updateUser, id)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

func (s *UserService) DeleteUser(id uint) error {
	return s.UserRepository.DeleteUser(id)
}
