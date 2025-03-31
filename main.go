package main

import (
	"log"
	"os"
	"url-shortener/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	routes.SetupRoutes(app)

	log.Println("Server is running on port 8080")
	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = ":3000"
	}

	log.Fatal(app.Listen(PORT))
}
