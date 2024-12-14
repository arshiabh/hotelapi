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

func main(){
	listenAddr := flag.String("listenAddr", "localhost:8080", "listeing to serve router")
	flag.Parse()

	app := fiber.New()
	apiv1 := app.Group("api/v1/")

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dburi))

	if err != nil {
		log.Fatal(err)
	}

	userhandler := api.NewUserHandler(db.NewMongoUserStore(client))
	apiv1.Get("user/", userhandler.HandleGetUser)
	apiv1.Get("user/:id", userhandler.HandleGetUsers)

	app.Listen(*listenAddr)
}




