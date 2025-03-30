package handlers

import (
	"strconv"

	"github.com/CaioProg/GoFinance/internal/models"
	"github.com/CaioProg/GoFinance/internal/services"
	"github.com/gofiber/fiber/v2"
)

type ExpenseHandler struct {
	ExpenseService services.ExpenseService
}

func (h *ExpenseHandler) GetExpenseById(ctx *fiber.Ctx) error {
	id, err := strconv.ParseUint(ctx.Params("id"), 10, 32)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Id"})
	}

	expense, err := h.ExpenseService.GetExpenseById(uint(id))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Expense not found"})
	}

	return ctx.Status(fiber.StatusOK).JSON(expense)
}

func (h *ExpenseHandler) GetAllExpenses(ctx *fiber.Ctx) error {
	expenses, err := h.ExpenseService.GetAllExpenses()
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Expenses not found"})
	}

	return ctx.Status(fiber.StatusOK).JSON(expenses)
}

func (h *ExpenseHandler) CreateExpense(ctx *fiber.Ctx) error {
	var expense models.Expense

	if err := ctx.BodyParser(&expense); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request data",
		})
	}

	expenseCreated, err := h.ExpenseService.CreateExpense(&expense)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(expenseCreated)
}

func (h *ExpenseHandler) UpdateExpense(ctx *fiber.Ctx) error {
	var expense models.Expense

	id, err := strconv.ParseUint(ctx.Params("id"), 10, 32)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Id"})
	}

	if err := ctx.BodyParser(&expense); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request data",
		})
	}

	expenseUpdated, err := h.ExpenseService.UpdateExpense(&expense, uint(id))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(expenseUpdated)
}

func (h *ExpenseHandler) DeleteExpense(ctx *fiber.Ctx) error {
	id, err := strconv.ParseUint(ctx.Params("id"), 10, 32)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid Id"})
	}

	err = h.ExpenseService.DeleteExpense(uint(id))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Expense deleted successfully"})
}
