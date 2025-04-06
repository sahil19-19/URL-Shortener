package main

import (
	"log"
	"os"
	"url-shortener/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	if os.Getenv("RENDER") == "" {
		// means it is being run in local
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading env variables")
		}
	}

	app := fiber.New()
	app.Use(logger.New())

	app.Use(cors.New(cors.Config{
		AllowOrigins: os.Getenv("FRONT_END_URL"), // Replace with your actual frontend URL
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	routes.SetupRoutes(app)

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = ":5000"
	}

	log.Println("Server is running on port:", PORT)

	log.Fatal(app.Listen(PORT))
}
