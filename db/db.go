package db

const (
	DBNAME     = "hotel-reservation"
	DBURI      = "mongodb://localhost:27017"
	TestDBURI  = DBURI
	TestDBNAME = "hotel-reservation-test"
)

type Pagination struct {
	Limit int64
	Page  int64
}

type Store struct {
	User    UserStore
	Hotel   HotelStore
	Room    RoomStore
	Booking BookingStore
}
