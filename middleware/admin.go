package middleware

import (
	"github.com/arshiabh/hotelapi/types"
	"github.com/gofiber/fiber/v2"
)

func AdminAuth(c *fiber.Ctx) error {
	user, ok := c.Context().UserValue("user").(*types.User)
	if !ok {
		return c.Status(400).JSON(fiber.Map{"error":"bad request"})
	}
	if !user.IsAdmin {
		return c.JSON(fiber.Map{"error":"not authorized"})
	}
	return c.Next()
}
