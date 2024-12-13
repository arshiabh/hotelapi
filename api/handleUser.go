package api

import (
	"github.com/arshiabh/hotelapi/types"
	"github.com/gofiber/fiber/v2"
)

func HandleGetUsers(c *fiber.Ctx) error{
	u := types.User{
		ID: "4",
		FirstName: "arina",
		LastName: "cheraghi",
	}

	return c.JSON(fiber.Map{"user":u})
}


func HandleGetUser(c *fiber.Ctx) error{
	id := c.Params("id")
	return c.JSON(fiber.Map{"userid":id})
}