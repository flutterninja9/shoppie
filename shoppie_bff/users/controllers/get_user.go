package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/dig"
)

func GetUser(c *fiber.Ctx, container *dig.Container) error {
	return c.SendStatus(200)

}
