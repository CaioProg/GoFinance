package repositories

import (
	"fmt"

	"github.com/CaioProg/GoFinance/internal/models"
	"gorm.io/gorm"
)

type ExpenseRepository interface {
	CreateExpense(expense *models.Expense) (*models.Expense, error)
	GetAllExpenses() (*[]models.Expense, error)
	DeleteExpense(id uint) error
	GetExpenseById(id uint) (*models.Expense, error)
	UpdateExpense(expense *models.Expense, id uint) (*models.Expense, error)
}

type ExpenseRepositoryImpl struct {
	DB *gorm.DB
}

func (r *ExpenseRepositoryImpl) GetExpenseById(id uint) (*models.Expense, error) {
	var expense models.Expense
	if err := r.DB.First(&expense, id).Error; err != nil {
		return nil, err
	}
	return &expense, nil
}

func (r *ExpenseRepositoryImpl) GetAllExpenses() (*[]models.Expense, error) {
	var expenses []models.Expense
	if err := r.DB.Find(&expenses).Error; err != nil {
		return nil, err
	}
	return &expenses, nil
}

func (r *ExpenseRepositoryImpl) CreateExpense(expense *models.Expense) (*models.Expense, error) {
	if err := r.DB.Create(&expense).Error; err != nil {
		return nil, err
	}
	return expense, nil
}

func (r *ExpenseRepositoryImpl) UpdateExpense(expense *models.Expense, id uint) (*models.Expense, error) {
	var existingExpense models.Expense
	if err := r.DB.First(&existingExpense, id).Error; err != nil {
		return nil, fmt.Errorf("expense with id %d not found", id)
	}

	if err := r.DB.Model(&existingExpense).Updates(expense).Error; err != nil {
		return nil, err
	}

	return &existingExpense, nil
}

func (r *ExpenseRepositoryImpl) DeleteExpense(id uint) error {
	var existingExpense models.Expense
	if err := r.DB.First(&existingExpense, id).Error; err != nil {
		return fmt.Errorf("expense with id %d not found", id)
	}

	if err := r.DB.Delete(&existingExpense).Error; err != nil {
		return err
	}

	return nil
}
