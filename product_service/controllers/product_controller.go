package controllers

import (
	"log"

	"github.com/flutterninja9/shoppie/product_service/database"
	"github.com/flutterninja9/shoppie/product_service/middlewares"
	"github.com/flutterninja9/shoppie/product_service/models"
	"github.com/gofiber/fiber/v2"
)

type UpdateProductRequest struct {
	Name        *string  `json:"name"`
	Description *string  `json:"description"`
	Price       *float64 `json:"price"`
	Quantity    *int     `json:"quantity"`
}

func GetAllProducts(c *fiber.Ctx) error {
	products, err := database.GetAllProducts()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Unable to fetch products",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"data": products,
	})
}

func GetProductById(c *fiber.Ctx) error {
	id := c.Params("id")

	product, err := database.GetProductById(id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "No such product found",
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"data": product,
	})
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
	id := c.Params("id")
	product, fetchErr := database.GetProductById(id)

	if fetchErr != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Product not found",
		})
	}

	updateReq := new(UpdateProductRequest)
	err := c.BodyParser(updateReq)
	if err != nil {
		return c.SendStatus(fiber.ErrBadRequest.Code)
	}

	if updateReq.Name != nil {
		product.Name = *updateReq.Name
	}
	if updateReq.Description != nil {
		product.Description = *updateReq.Description
	}
	if updateReq.Quantity != nil {
		product.Quantity = *updateReq.Quantity
	}
	if updateReq.Price != nil {
		product.Price = *updateReq.Price
	}

	updateErr := database.UpdateProduct(product)
	if updateErr != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Oops! something went wrong while saving",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Product updated successfully",
		"data":    product,
	})
}

func DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	log.Println("Trying to fetch product with id:", id)
	product, err := database.GetProductById(id)
	if err != nil {
		return c.SendStatus(500)
	}
	log.Println("Found product with id:", id)

	log.Println("Trying to delete product with id:", id)
	deleteErr := database.DeleteProduct(product)

	if deleteErr != nil {
		log.Println(deleteErr)
		return c.SendStatus(500)
	}

	log.Println("Deleted product with id:", id)
	return c.Status(200).JSON(fiber.Map{
		"message": "Product deleted",
	})
}
