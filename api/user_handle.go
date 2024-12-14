package api

import (
	"context"

	"github.com/arshiabh/hotelapi/db"
	"github.com/arshiabh/hotelapi/types"
	"github.com/gofiber/fiber/v2"
)

//must implement db.userstore
type UserHandler struct{
	UserStore db.UserStore
}


func NewUserHandler(userstore db.UserStore) *UserHandler{
	return &UserHandler{
		UserStore: userstore,
	}
}


func (h *UserHandler) HandleGetUser(c *fiber.Ctx) error{
	id := c.Params("id")
	user, err := h.UserStore.GetUserById(context.TODO(),id)
	if err != nil {
		return err
	}
	return c.JSON(user)
}


func (h *UserHandler) HandleGetUsers(c *fiber.Ctx) error{
	u := types.User{
		ID: "4",
		FirstName: "arina",
		LastName: "cheraghi",
	}

	return c.JSON(fiber.Map{"user":u})
}

