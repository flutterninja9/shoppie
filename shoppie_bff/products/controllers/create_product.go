package controllers

import (
	productsdk "github.com/flutterninja9/shoppie/product_sdk"
	middleware "github.com/flutterninja9/shoppie/shoppie_bff/middlewares"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/dig"
)

func CreateProduct(c *fiber.Ctx, container *dig.Container) error {
	authInfo := c.Locals("authInfo").(*middleware.AuthInfo)
	req := new(productsdk.CreateProductRequest)
	err := c.BodyParser(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
	}

	return container.Invoke(func(p productsdk.ProductSdk) error {
		product, err := p.CreateProduct(authInfo.Token, req)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Product created",
			"data":    product,
		})
	})
}
