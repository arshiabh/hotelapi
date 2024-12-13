package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/arshiabh/hotelapi/api"
	"github.com/arshiabh/hotelapi/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dburi = "mongodb://localhost:27017"
const dbname = "hotel-reservation"
const dbcoll = "users"


func main(){
	client, err := mongo.Connect(context.TODO(), options.Client().
	ApplyURI(dburi))

	if err != nil {
		log.Fatal(err)
	}

	user := types.User{
		FirstName: "arshia",
		LastName: "bohlooli",
	}
	ctx:= context.Background()
	
	coll := client.Database(dbname).Collection(dbcoll)
	res , err := coll.InsertOne(ctx, user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)

	listenAddr := flag.String("listenAddr", "localhost:8080", "listeing to serve router")
	flag.Parse()

	app := fiber.New()
	apiv1 := app.Group("api/v1/")

	apiv1.Get("user/", api.HandleGetUsers)
	apiv1.Get("user/:id", api.HandleGetUser)

	app.Listen(*listenAddr)
}




