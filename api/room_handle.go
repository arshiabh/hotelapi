package api

import (
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

}
