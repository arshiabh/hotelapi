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
	hotel := types.Hotel{
		Name:     "bellucia",
		Location: "france",
	}
	HotelStore.InsertHotel(context.TODO(), &hotel)

	room := types.Room{
		Type:      types.SingleRoomType,
		BasePrice: 99.9,
	}

	fmt.Println(hotel, room)
}
