package main

import (
	"log"
	"openguardian/routes"

	_ "openguardian/docs" // Import generated Swagger docs

	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

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

	// Register routes
	routes.SetupRoutes(app)

	// Start server
	port := "3000"
	log.Fatal(app.Listen(":" + port))
}
