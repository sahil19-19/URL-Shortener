package main

import (
	"log"
	"url-shortner/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	routes.SetupRoutes(app)

	log.Println("Server is running on port 8080")

	log.Fatal(app.Listen(":8080"))
}
