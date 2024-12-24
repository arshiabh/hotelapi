package types

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Booking struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserID   primitive.ObjectID `json:"userID" bson:"userID,omitempty"`
	RoomID   primitive.ObjectID `json:"roomID" bson:"roomID,omitempty"`
	FromDate time.Time          `json:"fromDate" bson:"fromDate,omitempty"`
	TillDate time.Time          `json:"tillDate" bson:"tillDate,omitempty"`
}
