package services

import (
	"github.com/CaioProg/GoFinance/internal/models"
	"github.com/CaioProg/GoFinance/internal/repositories"
)

type UserService struct {
	Repository repositories.UserRepository
}

func (s *UserService) GetUserByID(id uint) (*models.User, error) {
	return s.Repository.FindByID(id)
}
