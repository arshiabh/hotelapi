package api

import (
	"context"

	"github.com/arshiabh/hotelapi/db"
	"github.com/arshiabh/hotelapi/types"
	"github.com/gofiber/fiber/v2"
)

// must implement db.userstore
type UserHandler struct {
	UserStore db.UserStore
}

// want userstore that implement this
func NewUserHandler(userstore db.UserStore) *UserHandler {
	return &UserHandler{
		UserStore: userstore,
	}
}

func (h *UserHandler) HandleGetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := h.UserStore.GetUserById(context.TODO(), id)
	if err != nil {
		return err
	}
	return c.JSON(user)
}

func (h *UserHandler) HandleGetUsers(c *fiber.Ctx) error {
	users, err := h.UserStore.GetUsers(c.Context())
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"user": users})
}

func (h *UserHandler) HandlePostUser(c *fiber.Ctx) error {
	var params *types.CreateUserFromParams
	if err := c.BodyParser(&params); err != nil {
		return err
	}

	user, err := types.NewUserFromParams(*params)
	if err != nil {
		return err
	}

	inserteduser, err := h.UserStore.InsertUser(c.Context(), user)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"inserted user": inserteduser})
}
