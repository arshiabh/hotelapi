package main

import (
	"context"
	"flag"
	"log"

	"github.com/arshiabh/hotelapi/api"
	"github.com/arshiabh/hotelapi/db"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dburi = "mongodb://localhost:27017"

var config = fiber.Config{
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		return c.JSON(map[string]string{"error": err.Error()})
	},
}

func main() {
	listenAddr := flag.String("listenAddr", "localhost:8080", "listeing to serve router")
	flag.Parse()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dburi))
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New(config)
	apiv1 := app.Group("api/v1/")

	userhandler := api.NewUserHandler(db.NewMongoUserStore(client))
	apiv1.Post("user/", userhandler.HandlePostUser)
	apiv1.Get("user/", userhandler.HandleGetUsers)
	apiv1.Get("user/:id", userhandler.HandleGetUser)
	apiv1.Delete("user/:id", userhandler.HandleDeleteUser)
	apiv1.Put("user/:id", userhandler.HandlePutUser)

	app.Listen(*listenAddr)
}
