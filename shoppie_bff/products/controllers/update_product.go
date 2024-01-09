package controllers

import (
	productsdk "github.com/flutterninja9/shoppie/product_sdk"
	middleware "github.com/flutterninja9/shoppie/shoppie_bff/middlewares"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"go.uber.org/dig"
)

func UpdateProduct(c *fiber.Ctx, container *dig.Container) error {
	return container.Invoke(func(l *logrus.Logger, p productsdk.ProductSdk) error {
		productId := c.Params("productId")
		l.Info("Trying to update product:", productId)
		authInfo := c.Locals("authInfo").(*middleware.AuthInfo)
		req := new(productsdk.UpdateProductRequest)

		err := c.BodyParser(req)

		if err != nil {
			l.Fatal("Bad request")
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
		}

		ok, err := p.UpdateProduct(productId, authInfo.Token, *req)
		if !ok {
			l.Fatal("Unable to update product", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err})
		}

		l.Println("Product updated")
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Product updated"})
	})
}
