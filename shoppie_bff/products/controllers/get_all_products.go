package controllers

import (
	productsdk "github.com/flutterninja9/shoppie/product_sdk"
	middleware "github.com/flutterninja9/shoppie/shoppie_bff/middlewares"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/dig"
)

func GetAllProducts(c *fiber.Ctx, container *dig.Container) error {
	authInfo := c.Locals("authInfo").(*middleware.AuthInfo)
	return container.Invoke(func(p productsdk.ProductSdk) error {
		products, err := p.GetAllProducts(authInfo.Token)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err})
		}

		return c.Status(fiber.StatusOK).JSON(products)
	})
}
