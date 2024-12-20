package db

import (
	"context"
	"fmt"

	"github.com/arshiabh/hotelapi/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const usercoll = "users"

type Dropper interface {
	Drop(context.Context) error
}

type UserStore interface {
	Dropper

	GetUserById(context.Context, primitive.ObjectID) (*types.User, error)
	GetUsers(context.Context) ([]*types.User, error)
	InsertUser(context.Context, *types.User) (*types.User, error)
	DropUser(context.Context, primitive.ObjectID) error
	UpdateUser(context.Context, primitive.ObjectID, bson.M) error
}

type MongoUserStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoUserStore(client *mongo.Client, dbname string) *MongoUserStore {
	return &MongoUserStore{
		client: client,
		coll:   client.Database(dbname).Collection(usercoll),
	}
}

func (s *MongoUserStore) Drop(ctx context.Context) error {
	fmt.Println("---dropping user collection")
	return s.coll.Drop(ctx)
}

func (s *MongoUserStore) GetUserById(ctx context.Context, id primitive.ObjectID) (*types.User, error) {
	user := types.User{}
	if err := s.coll.FindOne(ctx, bson.M{"_id": id}).Decode(&user); err != nil {
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

func (s *MongoUserStore) DropUser(ctx context.Context, uid primitive.ObjectID) error {
	_, err := s.coll.DeleteOne(ctx, bson.M{"_id": uid})
	if err != nil {
		return err
	}
	return nil
}

func (s *MongoUserStore) UpdateUser(ctx context.Context, uid primitive.ObjectID, update bson.M) error {
	_, err := s.coll.UpdateOne(ctx, bson.M{"_id": uid},
		bson.M{
			"$set": update,
		},
	)
	if err != nil {
		return err
	}
	return nil
}
