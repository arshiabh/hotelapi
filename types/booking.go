package types

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CustomTime struct {
	time.Time
}

type Booking struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserID    primitive.ObjectID `json:"userID" bson:"userID,omitempty"`
	RoomID    primitive.ObjectID `json:"roomID" bson:"roomID,omitempty"`
	NumPerson int                `json:"numPerson" bson:"numPerson"`
	FromDate  time.Time          `json:"fromDate" bson:"fromDate,omitempty"`
	TillDate  time.Time          `json:"tillDate" bson:"tillDate,omitempty"`
}

type BookingParams struct {
	NumPerson int       `json:"numPerson"`
	FromDate  time.Time `json:"fromDate"`
	TillDate  time.Time `json:"tillDate"`
}


func (p *BookingParams) Validate() error {
	if time.Now().After(p.FromDate) || time.Now().After(p.TillDate) {
		return fmt.Errorf("could not book room from past")
	}
	if p.NumPerson > 4 || p.NumPerson ==0 {
		return fmt.Errorf("invalid number of person")
	}
	return nil
}