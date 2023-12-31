package productsdk

import (
	"encoding/json"
	"io"
)

type GetProductByIdResponse struct {
	Data *ProductEntity `json:"data"`
}

func GetProductByIdResponseFromJson(r io.Reader) (*GetProductByIdResponse, error) {
	product := new(GetProductByIdResponse)
	err := json.NewDecoder(r).Decode(product)

	return product, err
}
