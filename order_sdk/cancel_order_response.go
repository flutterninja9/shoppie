package ordersdk

import (
	"encoding/json"
	"io"
)

type CancelOrderResponse struct {
	Message string `json:"message"`
}

func CancelOrderResponseToJson(r io.Reader) (*CancelOrderResponse, error) {
	res := new(CancelOrderResponse)
	err := json.NewDecoder(r).Decode(res)

	if err != nil {
		return nil, err
	}

	return res, nil
}
