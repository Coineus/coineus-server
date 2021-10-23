package api

import (
	"encoding/json"
	"log"

	"github.com/coineus/coineus-server/internal/app"
	"github.com/coineus/coineus-server/internal/storage"
	"github.com/coineus/coineus-server/model"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v4"
)

func GetArchivedOperationByIdHandler(s storage.ArchivedOperationStore) fiber.Handler {
	return func(c *fiber.Ctx) error {
		operationId := c.Params("id")
		user := app.GetUserClaims(c)

		operation, err := s.GetArchivedOperationById(user.Id, operationId)
		if err != nil {
			log.Println(err)
			if err == pgx.ErrNoRows {
				return c.JSON(model.Response{
					Success: false,
					Error:   err.Error(),
					Data:    nil,
				})
			}

			return c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.JSON(model.Response{
			Success: true,
			Error:   "",
			Data:    operation,
		})
	}
}

func GetAllArchivedOperationsByUserIdHandler(s storage.ArchivedOperationStore) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := app.GetUserClaims(c)

		operations, err := s.GetAllArchivedOperations(user.Id)
		if err != nil {
			log.Println(err)
			if err == pgx.ErrNoRows {
				return c.JSON(model.Response{
					Success: false,
					Error:   err.Error(),
					Data:    nil,
				})
			}
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.JSON(model.Response{
			Success: true,
			Error:   "",
			Data:    operations,
		})
	}
}

func AddArchivedOperationHandler(s storage.ArchivedOperationStore) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := app.GetUserClaims(c)

		var operation model.ArchivedOperation
		err := json.Unmarshal(c.Body(), &operation)
		if err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		if operation.UserId != user.Id {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		err = s.AddArchivedOperation(operation)
		if err != nil {
			log.Println(err)
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.JSON(model.Response{
			Success: true,
			Error:   "",
		})
	}
}

func DeleteArchivedOperationHandler(s storage.ArchivedOperationStore) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := app.GetUserClaims(c)

		var operation model.ArchivedOperation
		err := json.Unmarshal(c.Body(), &operation)
		if err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		if operation.UserId != user.Id {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		err = s.DeleteArchivedOperation(operation)
		if err != nil {
			log.Println(err)
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.JSON(model.Response{
			Success: true,
			Error:   "",
		})
	}
}
