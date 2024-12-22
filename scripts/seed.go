package main

import (
	"context"
	"fmt"
	"log"

	"github.com/arshiabh/hotelapi/db"
	"github.com/arshiabh/hotelapi/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client     *mongo.Client
	RoomStore  *db.MongoRoomStore
	HotelStore *db.MongoHotelStore
)

func seedHotel(name string, location string, rating int) {
	hotel := types.Hotel{
		Name:     name,
		Location: location,
		Rooms:    []primitive.ObjectID{},
		Rating:   rating,
	}
	insertedhotel, err := HotelStore.InsertHotel(context.TODO(), &hotel)
	if err != nil {
		log.Fatal(err)
	}
	rooms := []types.Room{
		{
			Size: "small",
			Price: 99.9,
		},
		{
			Size: "medium",
			Price: 122.0,
		},
		{
			Size: "kingsize",
			Price: 139.99,
		},
	}

	for _, room := range rooms {
		room.HotelID = insertedhotel.ID
		insertedroom, err := RoomStore.InsertRoom(context.TODO(), &room)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(insertedroom)
	}

}

func main() {
	seedHotel("bellucia", "france", 3)
	seedHotel("cozy one", "netherland", 5)
	seedHotel("totenham", "london", 1)
}

func init() {
	var err error
	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(db.DBURI))
	if err != nil {
		log.Fatal(err)
	}
	//drop table here to make test clear
	client.Database(db.DBNAME).Collection("hotels").Drop(context.TODO())
	client.Database(db.DBNAME).Collection("rooms").Drop(context.TODO())

	HotelStore = db.NewMongoHotelStore(client)
	RoomStore = db.NewMongoRoomStore(client, HotelStore)

}
