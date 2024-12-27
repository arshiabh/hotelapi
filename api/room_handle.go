package api

import (
	"github.com/arshiabh/hotelapi/db"
	"github.com/arshiabh/hotelapi/types"
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
	var params types.BookingParams
	if err := c.BodyParser(&params); err != nil {
		return err
	}
	user, ok := c.Context().Value("user").(*types.User)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "internal server error"})
	}

	book := types.Booking{
		UserID:    user.ID,
		RoomID:    room.ID,
		NumPerson: params.NumPerson,
		FromDate:  params.FromDate,
		TillDate:  params.TillDate,
	}
	insertedbook, err := h.Store.Book.InsertBooking(c.Context(), &book)
	if err != nil {
		return err
	}
	return c.JSON(fiber.Map{"booked": insertedbook})
}
