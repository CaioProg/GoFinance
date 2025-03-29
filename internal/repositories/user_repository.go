package repositories

import (
	"github.com/CaioProg/GoFinance/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindByID(id uint) (*models.User, error)
}

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func (r *UserRepositoryImpl) FindByID(id uint) (*models.User, error) {
	var user models.User
	if err := r.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
