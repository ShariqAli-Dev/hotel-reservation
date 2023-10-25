package main

import (
	"context"
	"fmt"
	"log"

	"github.com/shariqali-dev/hotel-reservation/db"
	"github.com/shariqali-dev/hotel-reservation/types"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx := context.Background()
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.DBURI))
	if err != nil {
		log.Fatal(err)
	}
	hotelStore := db.NewMongoHotelStore(client, db.DBNAME)
	roomStore := db.NewMongoRoomStore(client, db.DBNAME)

	hotel := types.Hotel{
		Name:     "Hotella",
		Location: "New York",
	}
	rooms := []types.Room{{
		Type:      types.SinglePersonRoomType,
		BasePrice: 99.99,
	}, {
		Type:      types.DeluxRoomType,
		BasePrice: 199.99,
	}, {
		Type:      types.SeaSideRoomType,
		BasePrice: 122.99,
	}}
	insertedHotel, err := hotelStore.InsertHotel(ctx, &hotel)
	if err != nil {
		log.Fatal(err)
	}
	for _, room := range rooms {
		room.HotelID = insertedHotel.ID
		insertedRoom, err := roomStore.InsertRoom(ctx, &room)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(insertedRoom)
	}
}
