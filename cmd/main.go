package main

import (
	"fmt"
	"log"

	"github.com/CaioProg/GoFinance/internal/db"
	"github.com/CaioProg/GoFinance/internal/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	db.Connect()
	//db.Migrate(db.DB)

	routes.AddRoutes(app, db.DB)

	err := app.Listen(":8080")
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}

	fmt.Println("Server is running on http://localhost:8080")
}
