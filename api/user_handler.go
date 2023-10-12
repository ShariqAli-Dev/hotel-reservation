package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shariqali-dev/hotel-reservation/types"
)

func HandleGetUsers(c *fiber.Ctx) error {
	user := types.User{
		FirstName: "Shariq",
		LastName:  "Ali",
	}
	return c.JSON(user)
}

func HandleGetUser(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"name": "james"})
}
