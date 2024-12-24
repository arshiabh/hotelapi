package main

import (
	"context"
	"flag"
	"log"

	"github.com/arshiabh/hotelapi/api"
	"github.com/arshiabh/hotelapi/db"
	"github.com/arshiabh/hotelapi/middleware"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var config = fiber.Config{
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		return c.JSON(map[string]string{"error": err.Error()})
	},
}

func main() {
	listenAddr := flag.String("listenAddr", "localhost:8080", "listeing to serve router")
	flag.Parse()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.DBURI))
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New(config)
	apiv1 := app.Group("api/v1/", middleware.JWTAuthentication)
	auth := app.Group("api/", )

	var (
		hotelstore = db.NewMongoHotelStore(client)
		roomstore  = db.NewMongoRoomStore(client, hotelstore)
		userstore  = db.NewMongoUserStore(client)
		store  = db.Store{
			User: userstore,
			Room: roomstore,
			Hotel: hotelstore,
		}
	)
	
	authhandler := api.NewAuthHandler(userstore)
	auth.Post("auth/", authhandler.HandleAuthenticate)
	
	userhandler := api.NewUserHandler(userstore)
	apiv1.Post("user/", userhandler.HandlePostUser)
	apiv1.Get("user/", userhandler.HandleGetUsers)
	apiv1.Get("user/:id", userhandler.HandleGetUser)
	apiv1.Delete("user/:id", userhandler.HandleDeleteUser)
	apiv1.Put("user/:id", userhandler.HandlePutUser)

	hotelhandler := api.NewHotelHandler(store)
	apiv1.Get("hotel/:id", hotelhandler.HandleGetHotel)
	apiv1.Get("hotel/:id/rooms/", hotelhandler.HandleGetHotelRooms)
	apiv1.Get("hotel/", hotelhandler.HandleGetHotels)
	
	roomhandler := api.NewRoomHandler(store)
	apiv1.Get("room/:id/book/", roomhandler.HandleBookRoom)


	app.Listen(*listenAddr)
}
