package db

import (
	"context"

	"github.com/arshiabh/hotelapi/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var bookingcoll = "booking"

type BookingStore interface {
	InsertBooking(context.Context, *types.Booking) (*types.Booking, error)
}

type MongoBookStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoBookingStore(client *mongo.Client) *MongoBookStore {
	return &MongoBookStore{
		client: client,
		coll:   client.Database(DBNAME).Collection(bookingcoll),
	}
}

func (s *MongoBookStore) InsertBooking(ctx context.Context, data *types.Booking) (*types.Booking, error) {
	res, err := s.coll.InsertOne(ctx, data)
	if err != nil {
		return nil, err
	}
	data.ID = res.InsertedID.(primitive.ObjectID)
	return data, nil
}
