package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func GetCart(c *fiber.Ctx) error {
	return c.SendStatus(200)
}
