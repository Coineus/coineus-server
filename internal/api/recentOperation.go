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

func GetOperationByIdHandler(s storage.RecentOperationStore) fiber.Handler {
	return func(c *fiber.Ctx) error {
		operationId := c.Params("id")
		user := app.GetUserClaims(c)

		operation, err := s.GetRecentOperationById(user.Id, operationId)
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

func GetAllOperationsByUserIdHandler(s storage.RecentOperationStore) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := app.GetUserClaims(c)

		operations, err := s.GetAllRecentOperations(user.Id)
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

func AddOperationHandler(s storage.RecentOperationStore) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := app.GetUserClaims(c)

		var operation model.RecentOperation
		err := json.Unmarshal(c.Body(), &operation)
		if err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		if operation.UserId != user.Id {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		if op, err := s.GetRecentOperationByCoinSymbol(operation.CoinSymbol); err != nil && op.CoinSymbol == operation.CoinSymbol {
			operation.BuyCost = ((op.BuyCost * op.CoinAmount) + operation.BuyCost*operation.CoinAmount) / operation.CoinAmount
			operation.CoinAmount = op.CoinAmount + operation.CoinAmount

			err = s.UpdateRecentOperation(operation)
			if err != nil {
				log.Println(err)
				return c.SendStatus(fiber.StatusInternalServerError)
			}
		} else {
			err = s.AddRecentOperation(operation)
			if err != nil {
				log.Println(err)
				return c.SendStatus(fiber.StatusInternalServerError)
			}
		}

		err = app.SendNewPairToPriceService(operation.CoinSymbol)
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

func DeleteOperationHandler(s storage.RecentOperationStore) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := app.GetUserClaims(c)

		var operation model.RecentOperation
		err := json.Unmarshal(c.Body(), &operation)
		if err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		if operation.UserId != user.Id {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		err = s.DeleteRecentOperation(operation)
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

func UpdateOperationHandler(s storage.RecentOperationStore) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := app.GetUserClaims(c)

		var operation model.RecentOperation
		err := json.Unmarshal(c.Body(), &operation)
		if err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		if operation.UserId != user.Id {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		err = s.UpdateRecentOperation(operation)
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
