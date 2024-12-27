package db

const (
	DBURI  = "mongodb://localhost:27017"
	DBNAME = "hotel-reservation"
)

type Store struct {
	User  UserStore
	Room  Roomstore
	Hotel Hotelstore
	Book BookingStore
}
