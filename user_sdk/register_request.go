package usersdk

import (
	"bytes"
	"encoding/json"
	"io"
)

type RegisterRequest struct {
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (r *RegisterRequest) ToJson() io.Reader {
	buffer := new(bytes.Buffer)
	json.NewEncoder(buffer).Encode(r)

	return buffer
}
