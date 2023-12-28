package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func RemoveFromCart(c *fiber.Ctx) error {
	return c.SendStatus(200)
}
