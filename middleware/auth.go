package middleware

import (
	"strings"

	"github.com/arshiabh/hotelapi/db"
	"github.com/arshiabh/hotelapi/utils"
	"github.com/gofiber/fiber/v2"
)

const Secret_key = "supersecret"

func JWTAuthentication(userstore db.UserStore) fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.GetReqHeaders()["X-Api-Token"]
		tokenstr := strings.Join(token, "\n")
		if tokenstr == "" {
			return c.Status(401).JSON(fiber.Map{"error": "unauthorized"})
		}
		//get item from valid token and set it to response header 
		userID, email, err := utils.VerifyToken(tokenstr)
		if err != nil {
			c.Status(400)
			return err
		}
		user, err := userstore.GetUserById(c.Context(),userID)
		if err != nil {
			return err
		}
		c.Context().SetUserValue("user", user)
		c.Set("userID", userID)
		c.Set("email", email)
		return c.Next()
	}
}
