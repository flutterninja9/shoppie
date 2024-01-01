package ordersdk

import (
	"bytes"
	"encoding/json"
	"io"
)

type UpdateOrderStatusRequest struct {
	Status string `json:"status" validate:"oneof=placed delivered outForDelivery cancelled inProcessing onHold returned"`
}

func (u *UpdateOrderStatusRequest) ToJson() (io.Reader, error) {
	buffer := new(bytes.Buffer)
	err := json.NewEncoder(buffer).Encode(u)

	if err != nil {
		return nil, err
	}
	return buffer, nil
}
