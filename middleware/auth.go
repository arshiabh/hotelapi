package middleware

import (

	"github.com/gofiber/fiber/v2"
)

func JWTauthentication(c *fiber.Ctx) error {
	token, ok := c.GetReqHeaders()["X-Api-Token"]
	if !ok {
		return c.Status(400).JSON(fiber.Map{"error":"unauthorized"})
	}
	return c.JSON(fiber.Map{"token":token})
}