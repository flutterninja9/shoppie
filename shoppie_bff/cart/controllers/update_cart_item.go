package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func UpdateCartItem(c *fiber.Ctx) error {
	return c.SendStatus(200)
}
