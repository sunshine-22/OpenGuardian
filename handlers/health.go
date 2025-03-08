package handlers

import (
	"github.com/gofiber/fiber/v2"
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
