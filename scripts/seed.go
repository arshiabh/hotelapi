package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/arshiabh/hotelapi/db"
	fixtures "github.com/arshiabh/hotelapi/db/fixture"
	"github.com/arshiabh/hotelapi/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	RoomStore  *db.MongoRoomStore
	HotelStore *db.MongoHotelStore
)

func main() {
	var ctx context.Context
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(db.DBURI))
	if err != nil {
		log.Fatal(err)
	}
	if err := client.Database(db.DBNAME).Drop(ctx); err != nil {
		log.Fatal(err)
	}
	hotelStore := db.NewMongoHotelStore(client)
	store := &db.Store{
		User:  db.NewMongoUserStore(client),
		Book:  db.NewMongoBookingStore(client),
		Room:  db.NewMongoRoomStore(client, hotelStore),
		Hotel: hotelStore,
	}

	user := fixtures.AddUser(store, "arshia", "bohlooli", false)
	id := user.ID.Hex()
	fmt.Println("james ->", utils.GenarateToken(id, user.Email))
	admin := fixtures.AddUser(store, "admin", "admin", true)
	fmt.Println("admin ->", utils.GenarateToken(admin.ID.Hex(), admin.Email))
	hotel := fixtures.AddHotel(store, "some hotel", "bermuda", 5, nil)
	room := fixtures.AddRoom(store, "large", true, 88.44, hotel.ID)
	booking := fixtures.AddBooking(store, user.ID, room.ID, time.Now(), time.Now().AddDate(0, 0, 5))
	fmt.Println("booking ->", booking.ID)

}
