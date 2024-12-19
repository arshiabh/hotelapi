package db

import (
	"context"

	"github.com/arshiabh/hotelapi/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const collname = "hotels"

type Hotelstore interface {
	InsertHotel(context.Context, *types.Hotel) (*types.Hotel, error)
	UpdateHotel(context.Context, primitive.ObjectID, bson.M) error
}

type MongoHotelStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoHotelStore(client *mongo.Client) *MongoHotelStore {
	return &MongoHotelStore{
		client: client,
		coll:   client.Database(DBNAME).Collection(collname),
	}
}

func (s *MongoHotelStore) InsertHotel(ctx context.Context, hotel *types.Hotel) (*types.Hotel, error) {
	res, err := s.coll.InsertOne(ctx, hotel)
	if err != nil {
		return nil, err
	}
	hotel.ID = res.InsertedID.(primitive.ObjectID)
	return hotel, nil
}

func (s *MongoHotelStore) UpdateHotel(ctx context.Context, hid primitive.ObjectID, update bson.M) error {
	_,err := s.coll.UpdateOne(ctx, bson.M{"_id":hid},
	bson.M{
		"$set":update,
	},
	)
	if err != nil {
		return err
	}
	return nil
}
