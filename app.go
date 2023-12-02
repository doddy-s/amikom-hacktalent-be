package main

import (
	"amikom-hacktalent-be/database"
	"amikom-hacktalent-be/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDB()

    app := fiber.New()

	routes.SetUpRoutes(app)

    app.Listen(":3030")
}