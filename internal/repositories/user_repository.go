package repositories

import (
	"fmt"

	"github.com/CaioProg/GoFinance/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindById(id uint) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	GetAllUsers() (*[]models.User, error)
	CreateUser(user *models.User) (*models.User, error)
	UpdateUser(user *models.User, id uint) (*models.User, error)
	DeleteUser(id uint) error
}

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func (r *UserRepositoryImpl) FindById(id uint) (*models.User, error) {
	var user models.User
	if err := r.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepositoryImpl) FindByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.DB.Where("email = ?", email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepositoryImpl) GetAllUsers() (*[]models.User, error) {
	var users []models.User
	if err := r.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return &users, nil
}

func (r *UserRepositoryImpl) CreateUser(user *models.User) (*models.User, error) {
	if err := r.DB.Create(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepositoryImpl) UpdateUser(user *models.User, id uint) (*models.User, error) {
	var existingUser models.User
	if err := r.DB.First(&existingUser, id).Error; err != nil {
		return nil, fmt.Errorf("user with id %d not found", id)
	}

	if err := r.DB.Model(&existingUser).Updates(user).Error; err != nil {
		return nil, err
	}

	return &existingUser, nil
}

func (r *UserRepositoryImpl) DeleteUser(id uint) error {
	var existingUser models.User
	if err := r.DB.First(&existingUser, id).Error; err != nil {
		return fmt.Errorf("user with id %d not found", id)
	}

	if err := r.DB.Delete(&existingUser).Error; err != nil {
		return err
	}

	return nil
}
