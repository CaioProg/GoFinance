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
}

func UserRouter(app *fiber.App, db *gorm.DB) {
	userRepository := &repositories.UserRepositoryImpl{DB: db}
	userService := &services.UserService{Repository: userRepository}
	userHandler := &handlers.UserHandler{UserService: *userService}

	app.Get("/users", userHandler.GetAllUsers)
	app.Get("/user/:id", userHandler.GetUserById)
	app.Post("/user", userHandler.CreateUser)
	app.Patch("/user/:id", userHandler.UpdateUser)
	app.Delete("/user/:id", userHandler.DeleteUser)
}
