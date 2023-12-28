package usersdk

import (
	"bytes"
	"encoding/json"
	"io"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (l *LoginRequest) ToJson() io.Reader {
	buffer := new(bytes.Buffer)
	json.NewEncoder(buffer).Encode(l)

	return buffer

}
