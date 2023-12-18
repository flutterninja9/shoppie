package controllers

import (
	"product_service/database"
	"product_service/middlewares"
	"product_service/models"

	"github.com/gofiber/fiber/v2"
)

func GetAllProducts(c *fiber.Ctx) error {
	products, err := database.GetAllProducts()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Unable to fetch products",
		})
	}

	return c.Status(200).JSON(products)
}

func GetProductById(c *fiber.Ctx) error {
	id := c.Params("id")

	product, err := database.GetProductById(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "No such product found",
		})
	}
	return c.Status(200).JSON(product)
}

func CreateProduct(c *fiber.Ctx) error {
	authData, ok := c.Locals("authInfo").(*middlewares.AuthInfo)
	if !ok {
		return c.SendStatus(401)
	}

	createdByUser := authData.Claims["user_id"].(string)
	product := new(models.ProductModel)
	err := c.BodyParser(product)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Check if all fields are sent",
		})
	}

	product.CreatedBy = createdByUser
	result, saveErr := database.CreateProduct(product)
	if saveErr != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Oops! something went wrong",
		})
	}

	return c.Status(200).JSON(result)
}

func UpdateProduct(c *fiber.Ctx) error {
	return c.SendStatus(200)
}

func DeleteProduct(c *fiber.Ctx) error {
	return c.SendStatus(200)
}
