package main

import (
	"fmt"

	productsdk "github.com/flutterninja9/shoppie/product_sdk"
)

func main() {
	baseUrl := "http://localhost:8000/products/v1"
	token := "_token_"
	sdk := productsdk.GetProductSdkInstance(baseUrl)

	// Get all products
	res, err := sdk.GetAllProducts(token)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res)

	// Create Product
	p := productsdk.CreateProductRequest{
		Name:        "Sample 2",
		Description: "Sample 2",
		Quantity:    100,
		Price:       1000,
	}

	createdProd, createErr := sdk.CreateProduct(token, &p)
	if createErr != nil {
		fmt.Println(createErr)
	}

	fmt.Println("Product created with Id:" + createdProd.ID.String())

	// Delete product
	_, delErr := sdk.DeleteProduct(token, "158a990e-92e1-45bf-80f4-aafd3612fcbc")
	if delErr != nil {
		fmt.Println(delErr)
	}
	fmt.Println("Product deleted")
}
