package routes

import (
	"openguardian/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/health", handlers.HealthCheck)
}
