package api

import (
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
		return err
	}
	if err := h.UserStore.Validation(c.Context(), Authparams); err != nil {
		c.Status(400) 
		return err 
	}
	token, err := utils.GenarateToken(Authparams.Email)
	if err != nil {
		return err
	}
	return c.Status(200).JSON(fiber.Map{"token": token})
}
