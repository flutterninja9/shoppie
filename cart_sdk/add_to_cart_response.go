package cartsdk

import (
	"encoding/json"
	"io"
)

type AddToCartRespose struct {
	Message string `json:"message"`
}

func AddToCartResposeFromJson(r io.Reader) (*AddToCartRespose, error) {
	res := new(AddToCartRespose)
	err := json.NewDecoder(r).Decode(res)

	return res, err
}
