package db

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	DBNAME     string
	DBURI      string
	TestDBURI  string
	TestDBNAME string
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

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	DBURI = os.Getenv("MONGO_DB_URL")
	DBNAME = os.Getenv("MONGO_DB_NAME")
	TestDBNAME = os.Getenv("TEST_DB_NAME")
	TestDBURI = os.Getenv("TEST_DB_URI")
}
