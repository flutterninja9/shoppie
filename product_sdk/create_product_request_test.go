package productsdk_test

import (
	"testing"

	productsdk "github.com/flutterninja9/shoppie/product_sdk"
)

func TestIsValid(t *testing.T) {
	p := new(productsdk.CreateProductRequest)
	p.Name = "Test"
	p.Description = "Test"
	p.Quantity = 2
	p.Price = 100.0

	// should pass as the struct is valid
	if !p.IsValid() {
		t.Error("Invalid product")
	}
}

func TestIsInValid(t *testing.T) {
	p := new(productsdk.CreateProductRequest)
	p.Name = "Test"
	p.Description = "Test"
	p.Quantity = 0
	p.Price = 100.0

	// should fail as the struct is valid
	if p.IsValid() {
		t.Error("This should have failed as this is a invalid product")
	}
}
