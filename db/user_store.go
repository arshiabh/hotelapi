package db

import (
	"context"

	"github.com/arshiabh/hotelapi/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const usercoll = "users"


type UserStore interface{
	GetUserById(context.Context, string)(*types.User, error)
}


type MongoUserStore struct{
	client *mongo.Client
	coll *mongo.Collection
}


func NewMongoUserStore(client *mongo.Client) *MongoUserStore{
	return &MongoUserStore{
		client: client,
		coll: client.Database(DBNAME).Collection(usercoll),
	}
}


func (s *MongoUserStore) GetUserById(ctx context.Context, id string)(*types.User, error){
	user := types.User{}
	if err := s.coll.FindOne(ctx, bson.M{"_id":ToObjectId(id)}).Decode(&user); err !=nil{
		return nil,err
	}
	return &user, nil
}