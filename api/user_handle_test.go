package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http/httptest"
	"testing"

	"github.com/arshiabh/hotelapi/db"
	"github.com/arshiabh/hotelapi/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type testdb struct {
	UserStore db.UserStore
}

func (tdb *testdb) teardown() {
	if err := tdb.UserStore.Drop(context.TODO()); err != nil {
		log.Fatal(err)
	}
}

func setup(t *testing.T) *testdb {
	// dbname := "hotel-reservation-test"
	testdburi := "mongodb://localhost:27017"
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(testdburi))
	if err != nil {
		log.Fatal(err)
	}
	t.Log()
	return &testdb{
		UserStore: db.NewMongoUserStore(client),
	}
}

func TestUserPost(t *testing.T) {
	tdb := setup(t)
	//build func to drop collection then defer it hear to delete it after test
	defer tdb.teardown()
	app := fiber.New()
	NewUserHandler := NewUserHandler(tdb.UserStore)
	app.Post("/", NewUserHandler.HandlePostUser)

	params := types.CreateUserFromParams{
		FirstName: "arshia",
		LastName:  "bohlooli",
		Email:     "arshitest@gmail.com",
		Password:  "1234",
	}
	
	b, _ := json.Marshal(params)
	req := httptest.NewRequest("POST", "/", bytes.NewReader(b))
	req.Header.Add("Content-Type", "application/json")
	res, _ := app.Test(req)
	fmt.Println(res.StatusCode)
}
