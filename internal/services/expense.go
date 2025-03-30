package services

import (
	"fmt"
	"time"

	"github.com/CaioProg/GoFinance/internal/models"
	"github.com/CaioProg/GoFinance/internal/repositories"
)

type ExpenseService struct {
	ExpenseRepository repositories.ExpenseRepository
	UserRepository    repositories.UserRepository
}

func (s *ExpenseService) GetExpenseById(id uint) (*models.Expense, error) {
	return s.ExpenseRepository.GetExpenseById(id)
}

func (s *ExpenseService) GetAllExpenses() (*[]models.Expense, error) {
	return s.ExpenseRepository.GetAllExpenses()
}

func (s *ExpenseService) CreateExpense(expense *models.Expense) (*models.Expense, error) {
	if expense.Date == "" {
		return nil, fmt.Errorf("date is required")
	}

	if expense.UserId == 0 {
		return nil, fmt.Errorf("user id is required")
	}

	existUser, err := s.UserRepository.GetUserById(uint(expense.UserId))
	if existUser == nil || err != nil {
		return nil, fmt.Errorf("user with id %d not found", expense.UserId)
	}

	//Validate CategoryId

	return s.ExpenseRepository.CreateExpense(expense)
}

func (s *ExpenseService) UpdateExpense(expense *models.Expense, id uint) (*models.Expense, error) {
	var UpdateExpense models.Expense
	UpdateExpense.UpdatedAt = time.Now()

	if expense.Date == "" {
		return nil, fmt.Errorf("date is required")
	}

	if expense.UserId == 0 {
		return nil, fmt.Errorf("user id is required")
	}

	existUser, err := s.UserRepository.GetUserById(uint(expense.UserId))
	if existUser == nil || err != nil {
		return nil, fmt.Errorf("user with id %d not found", expense.UserId)
	}

	UpdateExpense.UserId = expense.UserId

	if expense.Description != "" {
		UpdateExpense.Description = expense.Description
	}

	if expense.Title != "" {
		UpdateExpense.Title = expense.Title
	}

	if expense.CategoryId != 0 {
		UpdateExpense.CategoryId = expense.CategoryId
	}

	if expense.Date != "" {
		UpdateExpense.Date = expense.Date
	}

	updatedExpense, err := s.ExpenseRepository.UpdateExpense(&UpdateExpense, id)
	if err != nil {
		return nil, err
	}

	return updatedExpense, nil
}

func (s *ExpenseService) DeleteExpense(id uint) error {
	return s.ExpenseRepository.DeleteExpense(id)
}
