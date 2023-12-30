package productsdk

import (
	"encoding/json"
	"io"
)

type ProductsResult struct {
	Data []*ProductEntity `json:"data"`
}

func ProductsResultFromJson(r io.Reader) (*ProductsResult, error) {
	result := new(ProductsResult)
	err := json.NewDecoder(r).Decode(result)
	if err != nil {
		return nil, err
	}

	return result, err
}
