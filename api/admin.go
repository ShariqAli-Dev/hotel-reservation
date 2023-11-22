package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shariqali-dev/hotel-reservation/types"
)

func AdminAuth(c *fiber.Ctx) error {
	user, ok := c.Context().UserValue("user").(*types.User)
	if !ok {
		return ErrorUnauthorized()
	}
	if !user.IsAdmin {
		return ErrorUnauthorized()
	}
	return c.Next()
}
