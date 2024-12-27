package types

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CustomTime struct{
	time.Time
}

type Booking struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserID    primitive.ObjectID `json:"userID" bson:"userID,omitempty"`
	RoomID    primitive.ObjectID `json:"roomID" bson:"roomID,omitempty"`
	NumPerson int                `json:"numPerson" bson:"numPerson"`
	FromDate  string             `json:"fromDate" bson:"fromDate,omitempty"`
	TillDate  string             `json:"tillDate" bson:"tillDate,omitempty"`
}

type BookingParams struct {
	NumPerson int    `json:"numPerson"`
	FromDate  string `json:"fromDate"`
	TillDate  string `json:"tillDate"`
}

