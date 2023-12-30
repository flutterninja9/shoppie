package productsdk

import (
	"bytes"
	"encoding/json"
	"io"

	"github.com/go-playground/validator/v10"
)

type CreateProductRequest struct {
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Price       float64 `json:"price" validate:"required"`
	Quantity    int     `json:"quantity" validate:"required,gt=0"`
}

func (c *CreateProductRequest) ToJson() io.Reader {
	buffer := new(bytes.Buffer)
	json.NewEncoder(buffer).Encode(c)

	return buffer
}

func (c *CreateProductRequest) IsValid() bool {
	Validate = validator.New()
	err := Validate.Struct(c)

	return err == nil
}
