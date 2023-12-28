package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func GetAllProducts(c *fiber.Ctx) error {
	return c.SendStatus(200)
}
