package routes

import (
	"github.com/CaioProg/GoFinance/internal/handlers"
	"github.com/CaioProg/GoFinance/internal/repositories"
	"github.com/CaioProg/GoFinance/internal/services"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func AddRoutes(app *fiber.App, db *gorm.DB) {
	UserRouter(app, db)
	ExpenseRouter(app, db)
}

func UserRouter(app *fiber.App, db *gorm.DB) {
	userRepository := &repositories.UserRepositoryImpl{DB: db}
	userService := &services.UserService{UserRepository: userRepository}
	userHandler := &handlers.UserHandler{UserService: *userService}

	app.Get("/users", userHandler.GetAllUsers)
	app.Get("/user/:id", userHandler.GetUserById)
	app.Post("/user", userHandler.CreateUser)
	app.Patch("/user/:id", userHandler.UpdateUser)
	app.Delete("/user/:id", userHandler.DeleteUser)
}

func ExpenseRouter(app *fiber.App, db *gorm.DB) {
	expenseRepository := &repositories.ExpenseRepositoryImpl{DB: db}
	userRepository := &repositories.UserRepositoryImpl{DB: db}
	expenseService := &services.ExpenseService{ExpenseRepository: expenseRepository, UserRepository: userRepository}
	expenseHandler := &handlers.ExpenseHandler{ExpenseService: *expenseService}

	app.Get("/expenses", expenseHandler.GetAllExpenses)
	app.Get("/expense/:id", expenseHandler.GetExpenseById)
	app.Post("/expense", expenseHandler.CreateExpense)
	app.Patch("/expense/:id", expenseHandler.UpdateExpense)
	app.Delete("/expense/:id", expenseHandler.DeleteExpense)
}
