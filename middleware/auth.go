package middleware

import (
	"strings"

	"github.com/arshiabh/hotelapi/utils"
	"github.com/gofiber/fiber/v2"
)

const Secret_key = "supersecret"

func JWTAuthentication(c *fiber.Ctx) error {
	token := c.GetReqHeaders()["X-Api-Token"]
	tokenstr := strings.Join(token, "\n")
	if tokenstr == ""{
		return c.Status(401).JSON(fiber.Map{"error":"unauthorized"})
	}
	email, err := utils.VerifyToken(tokenstr)
	if err != nil {
		c.Status(400)
		return err
	}

	c.Set("email", email)
	return c.Next()
}

