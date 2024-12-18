package db

import (
	"context"

	"github.com/arshiabh/hotelapi/types"
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
}

func NewMongoRoomStore(client *mongo.Client, dbname string) *MongoRoomStore {
	return &MongoRoomStore{
		client: client,
		coll:   client.Database(dbname).Collection(coll),
	}
}

func (s *MongoRoomStore) InsertRoom(ctx context.Context, room *types.Room) (*types.Room, error) {
	res, err := s.coll.InsertOne(ctx, room)
	if err != nil {
		return nil, err
	}
	room.ID = res.InsertedID.(primitive.ObjectID)
	return room, nil
}


