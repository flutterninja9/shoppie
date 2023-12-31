package controllers

import (
	productsdk "github.com/flutterninja9/shoppie/product_sdk"
	middleware "github.com/flutterninja9/shoppie/shoppie_bff/middlewares"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"go.uber.org/dig"
)

func GetProductDetails(c *fiber.Ctx, container *dig.Container) error {
	productId := c.Params("productId")
	authInfo := c.Locals("authInfo").(*middleware.AuthInfo)

	return container.Invoke(func(l *logrus.Logger, p productsdk.ProductSdk) error {
		product, err := p.GetProductById(productId, authInfo.Token)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err})
		}

		return c.Status(fiber.StatusOK).JSON(product)
	})
}
