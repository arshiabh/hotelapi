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

func seedHotel(name string, location string) {
	hotel := types.Hotel{
		Name:     name,
		Location: location,
		Rooms:    []primitive.ObjectID{},
	}
	insertedhotel, err := HotelStore.InsertHotel(context.TODO(), &hotel)
	if err != nil {
		log.Fatal(err)
	}
	rooms := []types.Room{
		{
			Type:      types.SingleRoomType,
			BasePrice: 99.9,
		},
		{
			Type:      types.SeasideRoomType,
			BasePrice: 122.0,
		},
		{
			Type:      types.DeluxeRoomType,
			BasePrice: 139.99,
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

// func gethotel(hid string) {
// 	id, _ := primitive.ObjectIDFromHex(hid)
// 	hotel, _ := HotelStore.GetHotel(context.TODO(), id)
// 	fmt.Println(hotel)
// }


func main() {
	seedHotel("bellucia", "france")
	// gethotel("67642dcf64e7d66925170536")
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
