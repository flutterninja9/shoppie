package cartsdk

import (
	"bytes"
	"encoding/json"
	"io"
)

type UpdateCartItemRequest struct {
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

func (i *UpdateCartItemRequest) ToJson() (io.Reader, error) {
	bufffer := new(bytes.Buffer)

	err := json.NewEncoder(bufffer).Encode(i)
	return bufffer, err
}
