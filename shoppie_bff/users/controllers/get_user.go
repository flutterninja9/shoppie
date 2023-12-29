package controllers

import (
	"net/http"

	middleware "github.com/flutterninja9/shoppie/shoppie_bff/middlewares"
	usersdk "github.com/flutterninja9/shoppie/user_sdk"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"go.uber.org/dig"
)

func GetUser(c *fiber.Ctx, container *dig.Container) error {
	authInfo := c.Locals("authInfo").(*middleware.AuthInfo)

	return container.Invoke(func(l *logrus.Logger, u *usersdk.UserService) error {
		userId := authInfo.Claims["user_id"].(string)
		l.Info("Trying to get user with id: " + userId)
		user, err := u.GetUserById(userId, authInfo.Token)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err})

		}

		l.Info("Got user with id:" + userId)
		return c.Status(fiber.StatusOK).JSON(user)

	})

}
