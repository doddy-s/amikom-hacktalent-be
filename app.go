package main

import (
	"amikom-hacktalent-be/database"
	"amikom-hacktalent-be/routes"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDB()

	app := fiber.New()

	routes.SetUpRoutes(app)

	fmt.Println("GetJadwal Application")
	fmt.Println("Created by: doddy-s")
	fmt.Println("FullName: Mohamad Doddy Sujatmiko")
	fmt.Println("NIM: 21.11.4344")
	fmt.Println("Made with GOLANG, FIBER, and GORM")
	fmt.Println("AMIKOM HACKTALENT 2023")

	app.Listen(":3030")
}
