package api

import (

	"github.com/arshiabh/hotelapi/db"
	"github.com/arshiabh/hotelapi/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type HotelHandler struct {
	HotelStore db.Hotelstore
	RoomStore db.Roomstore
}

type HotelQueryparams struct {
	Rooms bool
	Rating int
}

func NewHotelHandler(hotelstore db.Hotelstore, roomstore db.Roomstore) *HotelHandler {
	return &HotelHandler{
		HotelStore: hotelstore,
		RoomStore: roomstore,
	}
}

func (h *HotelHandler) HandleGetHotels(c *fiber.Ctx) error {
	var qparams HotelQueryparams
	if err := c.QueryParser(&qparams); err != nil {
		return c.Status(400).JSON(fiber.Map{"error":"bad request"})
	}
	hotels, err := h.HotelStore.GetHotels(c.Context())
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "could not fetch data"})
	}
	return c.JSON(hotels)
}

func (h *HotelHandler) HandleGetHotel(c *fiber.Ctx) error {
	id := c.Params("id")
	hid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "bad request"})
	}
	hotel, err := h.HotelStore.GetHotelById(c.Context(), hid)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"message": "object not found"})
	}
	return c.JSON(fiber.Map{"hotel": hotel})
}


func (h *HotelHandler) HandleGetHotelRooms(c *fiber.Ctx) error {
	id := c.Params("id")
	hid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "bad request"})
	}
	rooms ,err := h.RoomStore.GetRooms(c.Context(), hid)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error":"could not fetch data"})
	}
	return c.JSON(fiber.Map{"rooms": rooms})
}


func (h *HotelHandler) HandlePostHotel(c *fiber.Ctx) error {
	var params types.Hotel
	if err := c.BodyParser(&params); err != nil {
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
	return c.JSON(fiber.Map{"inserted": insertedHotel})
}

func (h *HotelHandler) HandlePutHotel(c *fiber.Ctx) error {
	var params types.Hotel
	if err := c.BodyParser(&params); err != nil {
		return err
	}
	return nil
}
