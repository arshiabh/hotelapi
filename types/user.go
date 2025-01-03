package types

import (
	"fmt"
	"regexp"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type Authparams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUserFromParams struct {
	FirstName string ` json:"firstName" bson:"firstName"`
	LastName  string ` json:"lastName" bson:"lastName"`
}

func (p *UpdateUserFromParams) ToBson() bson.M {
	m := bson.M{}
	if len(p.FirstName) > 0 {
		m["firstName"] = p.FirstName
	}
	if len(p.LastName) > 0 {
		m["lastName"] = p.LastName
	}
	return m
}

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
	IsAdmin           bool               ` json:"IsAdmin" bson:"isAdmin"`
}

const (
	cost            = 14
	firstNamelength = 2
	lastNamelength  = 2
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

	if len(params.Password) < 3 {
		return fmt.Errorf("password is weak")
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
