package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/shariqali-dev/hotel-reservation/db"
	"github.com/shariqali-dev/hotel-reservation/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BookRoomParams struct {
	FromDate      time.Time `json:"fromDate"`
	TillDate      time.Time `json:"tillDate"`
	NumberPersons int       `json:"numberPersons"`
}

type RoomHandler struct {
	store *db.Store
}

func NewRoomHandler(store *db.Store) *RoomHandler {
	return &RoomHandler{
		store: store,
	}
}

func (h *RoomHandler) HandleBookRoom(c *fiber.Ctx) error {
	var params BookRoomParams
	if err := c.BodyParser(&params); err != nil {
		return err
	}
	roomOID, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return err
	}
	user, ok := c.Context().Value("user").(*types.User)
	if !ok {
		return c.Status(http.StatusInternalServerError).JSON(genericResp{
			Type: "error",
			Msg:  "internal server error",
		})
	}
	booking := types.Booking{
		UserID:        user.ID,
		RoomID:        roomOID,
		FromDate:      params.FromDate,
		TillDate:      params.TillDate,
		NumberPersons: params.NumberPersons,
	}
	fmt.Printf("%+v", booking)
	return nil
}
