package api

import (
	"log"

	"github.com/coineus/coineus-server/internal/app"
	"github.com/coineus/coineus-server/internal/storage"
	"github.com/gofiber/fiber/v2"
)

func HealtCheckHandler(s storage.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := app.GetUserClaims(c)
		log.Println("uname : ", user.UserName, " mail : ", user.Email)
		return c.SendStatus(200)
	}
}
