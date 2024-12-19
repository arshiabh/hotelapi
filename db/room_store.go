package db

import (
	"context"

	"github.com/arshiabh/hotelapi/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const coll = "rooms"

type Roomstore interface {
	InsertHotel(context.Context, *types.Hotel) (*types.Hotel, error)

}

type MongoRoomStore struct {
	client *mongo.Client
	coll   *mongo.Collection

	Hotelstore
}

func NewMongoRoomStore(client *mongo.Client, hotelstore Hotelstore) *MongoRoomStore {
	return &MongoRoomStore{
		client: client,
		coll:   client.Database(DBNAME).Collection(coll),
		Hotelstore: hotelstore,
	}
}

func (s *MongoRoomStore) InsertRoom(ctx context.Context, room *types.Room) (*types.Room, error) {
	res, err := s.coll.InsertOne(ctx, room)
	if err != nil {
		return nil, err
	}
	room.ID = res.InsertedID.(primitive.ObjectID)
	//update hotel here to put room id into hotel database
	update := bson.M{"$push":bson.M{"rooms":room.ID}}
	if err := s.Hotelstore.UpdateHotel(ctx, room.HotelID, update); err != nil {
		return nil, err
	}
	return room, nil
}


