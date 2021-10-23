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

func GetWalletByIdHandler(s storage.WalletStore) fiber.Handler {
	return func(c *fiber.Ctx) error {
		operationId := c.Params("id")
		user := app.GetUserClaims(c)

		wallet, err := s.GetWalletById(user.Id, operationId)
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
			Data:    wallet,
		})
	}
}

func GetAllWalletsByUserIdHandler(s storage.WalletStore) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := app.GetUserClaims(c)

		wallets, err := s.GetAllWallets(user.Id)
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
			Data:    wallets,
		})
	}
}

func AddWalletHandler(s storage.WalletStore) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := app.GetUserClaims(c)

		var wallet model.Wallet
		err := json.Unmarshal(c.Body(), &wallet)
		if err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		if wallet.UserId != user.Id {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		err = s.AddWallet(wallet)
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

func DeleteWalletHandler(s storage.WalletStore, walletOpStore storage.WalletOperationStore) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := app.GetUserClaims(c)

		var wallet model.Wallet
		err := json.Unmarshal(c.Body(), &wallet)
		if err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		if wallet.UserId != user.Id {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		err = s.DeleteWallet(wallet)
		if err != nil {
			log.Println(err)
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		err = walletOpStore.DeleteAllWalletOperations(wallet.Id)
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
