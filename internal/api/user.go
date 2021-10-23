package api

import (
	"encoding/json"
	"log"

	"github.com/coineus/coineus-server/internal/app"
	"github.com/coineus/coineus-server/internal/storage"
	"github.com/coineus/coineus-server/model"
	"github.com/gofiber/fiber/v2"
)

func GetUserHandler(s storage.UserStore) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userClaims := app.GetUserClaims(c)

		var user model.DBUser

		user, err := s.GetByMail(userClaims.Email)
		user.PasswordHash = ""
		if err != nil {
			log.Println(err)
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.JSON(model.Response{
			Success: true,
			Error:   "",
			Data:    user,
		})
	}
}

func UpdateUserHandler(s storage.UserStore) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userClaims := app.GetUserClaims(c)

		var user model.User
		err := json.Unmarshal(c.Body(), &user)
		if err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		if user.Id != userClaims.Id {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
		err = s.UpdateUser(user)
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

func DeleteUserHandler(s storage.UserStore) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userClaims := app.GetUserClaims(c)

		var user model.User
		err := json.Unmarshal(c.Body(), &user)
		if err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		if user.Id != userClaims.Id {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		err = s.DeleteUser(user)
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
