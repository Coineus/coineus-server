package api

import (
	"github.com/coineus/coineus-server/internal/storage"
	"github.com/gofiber/fiber/v2"
)

func LoginHandler(s storage.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// mail := c.FormValue("mail")
		// pass := c.FormValue("pass")
		return nil
	}
}

func RegisterHandler(s storage.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return nil
	}
}
