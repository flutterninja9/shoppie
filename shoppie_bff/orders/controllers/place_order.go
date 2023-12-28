package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func PlaceOrder(c *fiber.Ctx) error {
	return c.SendStatus(200)
}
