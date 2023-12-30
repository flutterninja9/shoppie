package productsdk

import (
	"encoding/json"
	"io"

	"github.com/google/uuid"
)

type ProductEntity struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Quantity    int       `json:"quantity"`
}

func CreateProductResponseFromJson(r io.Reader) (*ProductEntity, error) {
	res := new(ProductEntity)
	err := json.NewDecoder(r).Decode(res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
