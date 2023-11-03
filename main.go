package main

import (
	"context"
	"flag"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/shariqali-dev/hotel-reservation/api"
	"github.com/shariqali-dev/hotel-reservation/api/middleware"
	"github.com/shariqali-dev/hotel-reservation/db"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var config = fiber.Config{
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		return c.JSON(map[string]string{"error": err.Error()})
	},
}

func main() {
	// new 2023-11-03 18:52:09.2691161 -0400 EDT m=+0.004815001j
	// old 2023-11-03 18:51:41.900326 -0400 EDT m=+0.006637701
	listenAddr := flag.String("listenAddr", ":5000", "The listen address of the API server")
	flag.Parse()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.DBURI))
	if err != nil {
		log.Fatal(err)
	}
	// handlers init
	var (
		userStore  = db.NewMongoUserStore(client)
		hotelStore = db.NewMongoHotelStore(client)
		roomStore  = db.NewMongoRoomStore(client, hotelStore)
		store      = &db.Store{
			Hotel: hotelStore,
			User:  userStore,
			Room:  roomStore,
		}

		userHandler  = api.NewUserHandler(store.User)
		hotelHandler = api.NewHotelHandler(store)
		authHandler  = api.NewAuthHandler(store.User)
		roomHandler  = api.NewRoomHandler(store)

		app   = fiber.New(config)
		auth  = app.Group("/api")
		apiv1 = app.Group("/api/v1", middleware.JWTAuthentication(userStore))
	)

	// auth
	auth.Post("/auth", authHandler.HandleAuthenticate)

	// Versioned API Routes
	// user handlers
	apiv1.Put("/user/:id", userHandler.HandlePutUser)
	apiv1.Delete("/user/:id", userHandler.HandleDeleteUser)
	apiv1.Post("/user", userHandler.HandlePostUser)
	apiv1.Get("/user", userHandler.HandleGetUsers)
	apiv1.Get("/user/:id", userHandler.HandleGetUser)

	// hotel handlers
	apiv1.Get("/hotel", hotelHandler.HandleGetHotels)
	apiv1.Get("/hotel/:id", hotelHandler.HandleGetHotel)
	apiv1.Get("/hotel/:id/rooms", hotelHandler.HandleGetRooms)

	// room handlers
	apiv1.Post("/room/:id/book", roomHandler.HandleBookRoom)

	app.Listen(*listenAddr)
}
