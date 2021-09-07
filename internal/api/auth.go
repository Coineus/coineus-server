package api

import (
	"log"

	"github.com/coineus/coineus-server/internal/app"
	"github.com/coineus/coineus-server/internal/storage"
	"github.com/coineus/coineus-server/model"
	"github.com/gofiber/fiber/v2"
)

func LoginHandler(s storage.UserStore) fiber.Handler {
	return func(c *fiber.Ctx) error {
		mail := c.FormValue("email")
		pass := c.FormValue("password")

		dbUser, err := s.GetByMail(mail)
		if err != nil {
			log.Println(err)
		}
		if !app.CheckPasswordHash(pass, dbUser.PasswordHash) {
			return c.JSON(model.Response{
				Success: false,
				Error:   "WRONG_PASSWORD",
				Data:    nil,
			})
		}
		token, err := app.CreateToken(dbUser.Id, dbUser.Email, dbUser.UserName)
		if err != nil {
			return c.JSON(model.Response{
				Success: false,
				Error:   err.Error(),
			})
		}
		return c.JSON(model.Response{
			Success: true,
			Error:   "",
			Data: map[string]string{
				"token": token,
			},
		})
	}
}

func RegisterHandler(s storage.UserStore) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var userToRegister model.User
		userToRegister.Email = c.FormValue("email")
		userToRegister.Password = c.FormValue("password")
		userToRegister.UserName = c.FormValue("username")
		if !app.ValidateUser(userToRegister) {
			return c.JSON(model.Response{
				Success: false,
				Error:   "- Email must be in the correct format. For example : test@example.com\n- Username must be at least 4 characters \n- Password must be at least 8 characters",
			})
		}
		if err := s.AddUser(userToRegister); err != nil {
			return c.JSON(model.Response{
				Success: false,
				Error:   err.Error(),
			})
		}

		dbUser, err := s.GetByMail(userToRegister.Email)
		if err != nil {
			return c.JSON(model.Response{
				Success: false,
				Error:   err.Error(),
			})
		}

		token, err := app.CreateToken(dbUser.Id, dbUser.Email, dbUser.UserName)
		if err != nil {
			return c.JSON(model.Response{
				Success: false,
				Error:   err.Error(),
			})
		}
		return c.JSON(model.Response{
			Success: true,
			Error:   "",
			Data: map[string]string{
				"token": token,
			},
		})
	}
}
