package main

import (
	"log"

	"github.com/CaioProg/GoFinance/internal/db"
	"github.com/CaioProg/GoFinance/internal/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	db.Connect()
	db.Migrate(db.DB)

	routes.AddRoutes(app, db.DB)
	log.Fatal(app.Listen(":8080"))
}
