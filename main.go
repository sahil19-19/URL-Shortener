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

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = ":5000"
	}

	log.Println("Server is running on port:", PORT)

	log.Fatal(app.Listen(PORT))
}
