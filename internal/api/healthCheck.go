package api

import (
	"github.com/coineus/coineus-server/internal/storage"
	"github.com/gofiber/fiber/v2"
)

func HealtCheckHandler(s storage.Database) fiber.Handler {
	return func(c *fiber.Ctx) error {
		row := s.Db.QueryRow(c.Context(), "SELECT now()")
		var timestamp interface{}
		err := row.Scan(&timestamp)
		if err != nil || timestamp == nil {
			return c.SendStatus(500)
		}
		return c.SendStatus(200)
	}
}
