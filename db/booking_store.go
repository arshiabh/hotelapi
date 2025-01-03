package db

import (
	"context"
	"fmt"

	"github.com/arshiabh/hotelapi/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var bookingcoll = "booking"

type BookingStore interface {
	InsertBooking(context.Context, *types.Booking) (*types.Booking, error)
	GetBookings(context.Context, bson.M) ([]*types.Booking, error)
	GetBookingByID(context.Context, string) (*types.Booking, error)
	DeleteBooking(context.Context, bson.M) error
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

func (s *MongoBookStore) GetBookings(ctx context.Context, filter bson.M) ([]*types.Booking, error) {
	var bookings []*types.Booking
	res, err := s.coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	res.All(ctx, &bookings)
	return bookings, err
}

func (s *MongoBookStore) InsertBooking(ctx context.Context, data *types.Booking) (*types.Booking, error) {
	res, err := s.coll.InsertOne(ctx, data)
	if err != nil {
		return nil, err
	}
	data.ID = res.InsertedID.(primitive.ObjectID)
	return data, nil
}

func (s *MongoBookStore) GetBookingByID(ctx context.Context, id string) (*types.Booking, error) {
	var room *types.Booking
	bid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	res := s.coll.FindOne(ctx, bson.M{"_id": bid})
	if err := res.Decode(&room); err != nil {
		if err.Error() == mongo.ErrNoDocuments.Error() {
			return nil, fmt.Errorf("nothing found")
		}
		return nil, err
	}
	return room, nil
}

func (s *MongoBookStore) DeleteBooking(ctx context.Context, filter bson.M) error {
	_, err := s.coll.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}
