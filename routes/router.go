package routes

import (
	"url-shortener/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/api/health", controllers.HealthCheck)

	app.Post("/api/shorten", controllers.CreateShortURL)

	app.Get("/api/:shorturl", controllers.RedirectURL)
}
