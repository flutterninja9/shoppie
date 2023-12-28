package usersdk

import (
	"encoding/json"
	"io"
)

type LoginResponse struct {
	Token string `json:"token"`
}

func LoginResponseFromJson(b io.Reader) (*LoginResponse, error) {
	res := new(LoginResponse)
	err := json.NewDecoder(b).Decode(res)

	if err != nil {
		return nil, err
	}

	return res, nil
}
