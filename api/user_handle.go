package api

import (
	"github.com/arshiabh/hotelapi/db"
	"github.com/arshiabh/hotelapi/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	user, err := h.UserStore.GetUserById(c.Context(), id)
	if err != nil {
		return c.JSON(fiber.Map{"message": "data not found"})
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
	var params types.CreateUserFromParams
	if err := c.BodyParser(&params); err != nil {
		return err
	}

	if err := params.Validate(); err != nil {
		return err
	}

	user, err := types.NewUserFromParams(params)
	if err != nil {
		return err
	}

	inserteduser, err := h.UserStore.InsertUser(c.Context(), user)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"inserted user": inserteduser})
}

func (h *UserHandler) HandleDeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	uid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	if err := h.UserStore.DropUser(c.Context(), uid); err != nil {
		return err
	}
	return c.JSON(fiber.Map{"message": "user successfully removed"})
}

func (h *UserHandler) HandlePutUser(c *fiber.Ctx) error {
	id := c.Params("id")
	uid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	var params *types.UpdateUserFromParams
	if err := c.BodyParser(&params); err != nil {
		return err
	}

	update := params.ToBson()

	if err := h.UserStore.UpdateUser(c.Context(), uid, update); err != nil {
		return err
	}

	return c.JSON(fiber.Map{"message": "successfully updated!"})
}
