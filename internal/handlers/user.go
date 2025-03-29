package handlers

import (
	"strconv"

	"github.com/CaioProg/GoFinance/internal/services"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	UserService services.UserService
}

func (h *UserHandler) GetUserById(ctx *fiber.Ctx) error {
	id, err := strconv.ParseUint(ctx.Params("id"), 10, 32)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	user, err := h.UserService.GetUserByID(uint(id))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	return ctx.Status(fiber.StatusOK).JSON(user)
}
