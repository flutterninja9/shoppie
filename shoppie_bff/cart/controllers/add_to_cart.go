package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func AddToCart(c *fiber.Ctx, l *logrus.Logger) error {
	return c.SendStatus(200)
}
