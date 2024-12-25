package api

import (
	"fmt"

	"github.com/arshiabh/hotelapi/db"
	"github.com/gofiber/fiber/v2"
)

type RoomHandler struct {
	Store db.Store
}

func NewRoomHandler(store db.Store) *RoomHandler {
	return &RoomHandler{
		Store: store,
	}
}

func (h *RoomHandler) HandleBookRoom(c *fiber.Ctx) error {
	id := c.Params("id")
	room, err := h.Store.Room.GetRoomById(c.Context(), id)
	if err != nil {
		c.Status(fiber.StatusNotFound)
		return err
	}
	email := c.GetRespHeader("email")
	userID := c.GetRespHeader("userID")
	fmt.Println(email, userID)
	return c.JSON(fiber.Map{"room": room})
}
