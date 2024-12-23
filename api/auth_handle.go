package api

import (
	"fmt"

	"github.com/arshiabh/hotelapi/db"
	"github.com/arshiabh/hotelapi/middleware"
	"github.com/arshiabh/hotelapi/types"
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
		return err
	}
	if token, err := middleware.GenarateToken(Authparams.Email); err != nil {
		return err
		} else {
			fmt.Println(token)
		}
	return c.Status(200).JSON(fiber.Map{"message": "login successfully"})
}
