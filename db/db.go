package db

import "go.mongodb.org/mongo-driver/bson/primitive"

//for general db usage

const DBNAME = "hotel-reservation"

//for converting mongoid 
func ToObjectId(id string) primitive.ObjectID {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}

	return oid
}