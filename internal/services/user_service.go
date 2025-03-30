package services

import (
	"fmt"
	"time"

	"github.com/CaioProg/GoFinance/internal/models"
	"github.com/CaioProg/GoFinance/internal/repositories"
)

type UserService struct {
	Repository repositories.UserRepository
}

func (s *UserService) GetUserById(id uint) (*models.User, error) {
	return s.Repository.FindById(id)
}

func (s *UserService) GetAllUsers() (*[]models.User, error) {
	return s.Repository.GetAllUsers()
}

func (s *UserService) CreateUser(user *models.User) (*models.User, error) {

	existEmail, err := s.Repository.FindByEmail(user.Email)
	if err != nil {
		return nil, err
	}

	if existEmail != nil {
		return nil, fmt.Errorf("email '%s' already exists", user.Email)
	}

	return s.Repository.CreateUser(user)
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

	updatedUser, err := s.Repository.UpdateUser(&updateUser, id)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

func (s *UserService) DeleteUser(id uint) error {
	return s.Repository.DeleteUser(id)
}
