package api

import (
	"github.com/arshiabh/hotelapi/db"
	"github.com/arshiabh/hotelapi/types"
	"github.com/gofiber/fiber/v2"
)

type HotelHandler struct {
	HotelStore db.Hotelstore
}

func NewHotelHandler(hotelstore db.Hotelstore) *HotelHandler {
	return &HotelHandler{
		HotelStore: hotelstore,
	}
}

func (h *HotelHandler) HandlePostHotel(c *fiber.Ctx) error {
	var params types.Hotel
	if err := c.BodyParser(&params); err != nil{
		return err
	}
	hotel, err := h.HotelStore.InsertHotel(c.Context(), &params)
	if err != nil {
		return nil
	}
	insertedHotel, err := h.HotelStore.InsertHotel(c.Context(), hotel)
	if err != nil {
		return err
	}
	return c.JSON(fiber.Map{"inserted":insertedHotel})
}
