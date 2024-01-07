package controllers

import (
	productsdk "github.com/flutterninja9/shoppie/product_sdk"
	middleware "github.com/flutterninja9/shoppie/shoppie_bff/middlewares"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"go.uber.org/dig"
)

func UpdateProduct(c *fiber.Ctx, container *dig.Container) error {
	productId := c.Params("productId")
	authInfo := c.Locals("authInfo").(*middleware.AuthInfo)
	req := new(productsdk.UpdateProductRequest)

	err := c.BodyParser(req)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
	}

	return container.Invoke(func(l *logrus.Logger, p productsdk.ProductSdk) error {
		ok, err := p.UpdateProduct(productId, authInfo.Token, *req)
		if !ok {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Product updated"})
	})
}
