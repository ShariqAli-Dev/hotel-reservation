package db

const (
	DBNAME     = "hotel-reservation"
	DBURI      = "mongodb://localhost:27017"
	TestDBURI  = DBURI
	TestDBNAME = "hotel-reservation-test"
)

type Store struct {
	User  UserStore
	Hotel HotelStore
	Room  RoomStore
}
