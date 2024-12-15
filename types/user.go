package types

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

// create this just for POST request
type CreateUserFromParams struct {
	FirstName string ` json:"firstName"`
	LastName  string ` json:"lastName"`
	Email     string ` json:"email"`
	Password  string `json:"password"`
}

type User struct {
	ID                primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	FirstName         string             ` json:"firstName" bson:"firstName"`
	LastName          string             ` json:"lastName" bson:"lastName"`
	Email             string             ` json:"email" bson:"email"`
	EncryptedPassword string             ` json:"-" bson:"EncryptedPassword"`
}

func NewUserFromParams(params CreateUserFromParams) (*User, error) {
	EncryptedPassword, err := bcrypt.GenerateFromPassword([]byte(params.Password), 14)
	if err != nil {
		return nil, err
	}
	return &User{
		FirstName:         params.FirstName,
		LastName:          params.LastName,
		Email:             params.Email,
		EncryptedPassword: string(EncryptedPassword),
	}, nil
}
