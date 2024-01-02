package cartsdk

import (
	"encoding/json"
	"io"
)

type CommonCartRespose struct {
	Message string `json:"message"`
}

func CommonCartResposeFromJson(r io.Reader) (*CommonCartRespose, error) {
	res := new(CommonCartRespose)
	err := json.NewDecoder(r).Decode(res)

	return res, err
}
