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

func GetAllWalletOperationsByUserIdHandler(s storage.WalletOperationStore) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := app.GetUserClaims(c)

		walletOperationDTO, err := s.GetAllWalletOperation(user.Id)
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
			Data:    walletOperationDTO,
		})
	}
}

func AddWalletOperationHandler(s storage.WalletOperationStore) fiber.Handler {
	return func(c *fiber.Ctx) error {
		//user := app.GetUserClaims(c)

		var walletOperation model.WalletOperation
		err := json.Unmarshal(c.Body(), &walletOperation)
		if err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		err = s.AddWalletOperation(walletOperation)
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

func DeleteWalletOperationHandler(s storage.WalletOperationStore) fiber.Handler {
	return func(c *fiber.Ctx) error {
		//user := app.GetUserClaims(c)

		var walletOperation model.WalletOperation
		err := json.Unmarshal(c.Body(), &walletOperation)
		if err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		err = s.DeleteWalletOperation(walletOperation)
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
