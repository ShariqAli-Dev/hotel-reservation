package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/shariqali-dev/hotel-reservation/db/fixtures"
	"github.com/shariqali-dev/hotel-reservation/types"
)

func TestUserGetBooking(t *testing.T) {
	db := setup(t)
	defer db.teardown(t)

	var (
		user    = fixtures.AddUser(db.Store, "test", "name", false)
		hotel   = fixtures.AddHotel(db.Store, "bar hotel", "a", 4, nil)
		room    = fixtures.AddRoom(db.Store, "small", true, 4.4, hotel.ID)
		from    = time.Now()
		till    = from.AddDate(0, 0, 5)
		booking = fixtures.AddBooking(db.Store, user.ID, room.ID, from, till)

		app   = fiber.New()
		route = app.Group("/", JWTAuthentication(db.User))

		BookingHandler = NewBookingHandler(db.Store)
	)

	route.Get("/:id", BookingHandler.HandleGetBookings)
	req := httptest.NewRequest("GET", fmt.Sprintf("/%s", booking.ID.Hex()), nil)
	req.Header.Add("X-Api-Token", CreateTokenFromUser(user))
	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("non 200 code get %d", resp.StatusCode)
	}
	var bookingResp *types.Booking
	if err := json.NewDecoder(resp.Body).Decode(&bookingResp); err != nil {
		t.Fatal(err)
	}
	fmt.Println(bookingResp)
	if bookingResp.ID != booking.ID {
		t.Fatalf("expected %s got %s", booking.ID, bookingResp.ID)
	}
	if bookingResp.ID != booking.UserID {
		t.Fatalf("expected %s got %s", booking.UserID, bookingResp.UserID)
	}
}

func TestAdminGetBooking(t *testing.T) {
	db := setup(t)
	defer db.teardown(t)

	var (
		adminUser = fixtures.AddUser(db.Store, "hohohehe", "hohohehe", true)
		user      = fixtures.AddUser(db.Store, "test", "name", false)
		hotel     = fixtures.AddHotel(db.Store, "bar hotel", "a", 4, nil)
		room      = fixtures.AddRoom(db.Store, "small", true, 4.4, hotel.ID)
		from      = time.Now()
		till      = from.AddDate(0, 0, 5)
		booking   = fixtures.AddBooking(db.Store, user.ID, room.ID, from, till)

		app   = fiber.New()
		admin = app.Group("/", JWTAuthentication(db.User), AdminAuth)

		BookingHandler = NewBookingHandler(db.Store)
	)

	_ = booking
	admin.Get("/", BookingHandler.HandleGetBookings)
	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Add("X-Api-Token", CreateTokenFromUser(adminUser))
	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("non 200 response %d", resp.StatusCode)
	}
	var bookings []*types.Booking
	if err := json.NewDecoder(resp.Body).Decode(&bookings); err != nil {
		t.Fatal(err)
	}
	// if len(bookings) != 1 {
	// 	t.Fatalf("expected 1 booking but got %d", len(bookings))
	// }
	have := bookings[0]
	if have.ID != booking.ID {
		t.Fatalf("expected %s got %s", booking.ID, have.ID)
	}
	if have.UserID != booking.UserID {
		t.Fatalf("expected %s got %s", booking.UserID, have.UserID)
	}
	// test non-admin acnnot access the bookings
	req = httptest.NewRequest("GET", "/", nil)
	req.Header.Add("X-Api-Token", CreateTokenFromUser(user))
	resp, err = app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode == http.StatusOK {
		t.Fatalf("expected a non 200 status code got %d", resp.StatusCode)
	}
}
