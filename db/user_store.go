package db

import (
	"context"

	"github.com/arshiabh/hotelapi/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const usercoll = "users"

type UserStore interface {
	GetUserById(context.Context, string) (*types.User, error)
	GetUsers(context.Context) ([]*types.User, error)
	InsertUser(context.Context, *types.User) (*types.User, error)
	DropUser(context.Context, *types.User) (error)
}

type MongoUserStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoUserStore(client *mongo.Client) *MongoUserStore {
	return &MongoUserStore{
		client: client,
		coll:   client.Database(DBNAME).Collection(usercoll),
	}
}

func (s *MongoUserStore) GetUserById(ctx context.Context, id string) (*types.User, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	user := types.User{}
	if err := s.coll.FindOne(ctx, bson.M{"_id": oid}).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *MongoUserStore) GetUsers(ctx context.Context) ([]*types.User, error) {
	cur, err := s.coll.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	users := []*types.User{}
	if err := cur.All(ctx, &users); err != nil {
		return []*types.User{}, nil
	}
	return users, nil

}

func (s *MongoUserStore) InsertUser(ctx context.Context, user *types.User) (*types.User, error) {
	res, err := s.coll.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	//id is not fill so we add this manually
	user.ID = res.InsertedID.(primitive.ObjectID)

	return user, nil
}

func (s *MongoUserStore) DropUser(ctx context.Context, user *types.User) error {
	_, err := s.coll.DeleteOne(ctx, bson.M{"_id":user.ID})
	if err != nil {
		return err
	}
	return nil
}