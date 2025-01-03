package api

import (
	"errors"
	"fmt"

	"github.com/arshiabh/hotelapi/db"
	"github.com/arshiabh/hotelapi/types"
	"github.com/arshiabh/hotelapi/utils"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	UserStore db.UserStore
}

// want userstore that implement this
func NewAuthHandler(userstore db.UserStore) *AuthHandler {
	return &AuthHandler{
		UserStore: userstore,
	}
}

func (h *AuthHandler) HandleAuthenticate(c *fiber.Ctx) error {
	var Authparams types.Authparams
	if err := c.BodyParser(&Authparams); err != nil {
		c.Status(fiber.StatusBadRequest)
		if errors.Is(err, fiber.ErrUnprocessableEntity) {
			c.Status(fiber.StatusUnprocessableEntity)
			return fmt.Errorf("some required information is missing")
		}
		return err
	}
	//check credentials and get useid for token
	userID, err := h.UserStore.Validation(c.Context(), Authparams)
	if err != nil {
		c.Status(400)
		return err
	}
	struserID := userID.Hex()
	token := utils.GenarateToken(struserID, Authparams.Email)
	return c.Status(200).JSON(fiber.Map{"token": token})
}
