package types

import (
	"fmt"
	"regexp"

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

const (
	cost = 14
	firstNamelength = 2
	lastNamelength = 2

)

func (params CreateUserFromParams) Validate() error {
	if len(params.FirstName) < firstNamelength {
		return fmt.Errorf("firstname is too short")
	}

	if len(params.LastName) < lastNamelength {
		return fmt.Errorf("lastname is too short")
	}

	if !isEmailValid(params.Email) {
		return fmt.Errorf("email is invalid")
	}

	return nil
}

func isEmailValid(e string) bool {
    emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
    return emailRegex.MatchString(e)
}

func NewUserFromParams(params CreateUserFromParams) (*User, error) {
	EncryptedPassword, err := bcrypt.GenerateFromPassword([]byte(params.Password), cost)
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
