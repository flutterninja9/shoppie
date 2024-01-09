package productsdk

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
)

type UpdateProductRequest struct {
	Name        *string  `json:"name"`
	Description *string  `json:"description"`
	Price       *float64 `json:"price"`
	Quantity    *int     `json:"quantity" validate:"gt=0"`
}

func (u *UpdateProductRequest) ToJson() io.Reader {
	buffer := new(bytes.Buffer)
	log.Println(u)
	json.NewEncoder(buffer).Encode(u)

	return buffer
}
