package main

import (
	"log"

	_ "openguardian/docs" // Import generated Swagger docs

	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

// HealthCheck godoc
// @Summary Health check endpoint
// @Description Returns server status
// @Tags health
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string
// @Router /api/health [get]
func HealthCheck(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "ok"})
}

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	app := fiber.New()

	cfg := swagger.Config{
		BasePath: "/",
		FilePath: "./docs/swagger.json",
		Path:     "/docs",
		Title:    "OpenGuradian Api Docs",
	}

	app.Use(swagger.New(cfg))

	api := app.Group("/api")
	api.Get("/health", HealthCheck)

	// Start server
	port := "3000"
	log.Fatal(app.Listen(":" + port))
}
