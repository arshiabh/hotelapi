package main

import (
	"context"
	"fmt"
	"log"

	"github.com/arshiabh/hotelapi/db"
	"github.com/arshiabh/hotelapi/types"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.DBURI))
	if err != nil {
		log.Fatal(err)
	}
	HotelStore := db.NewMongoHotelStore(client, db.DBNAME)
	RoomStore := db.NewMongoRoomStore(client, db.DBNAME)

	hotel := types.Hotel{
		Name:     "bellucia",
		Location: "france",
	}
	insertedhotel, err := HotelStore.InsertHotel(context.TODO(), &hotel)
	if err != nil {
		log.Fatal(err)
	}

	room := types.Room{
		Type:      types.SingleRoomType,
		BasePrice: 99.9,
	}
	
	room.HotelID = insertedhotel.ID
	insertedroom , err := RoomStore.InsertRoom(context.TODO(), &room)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(insertedhotel)
	fmt.Println(insertedroom)
}
