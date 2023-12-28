package usersdk

import (
	"bytes"
	"encoding/json"
	"io"

	"github.com/google/uuid"
)

type UserEntity struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
}

func (u *UserEntity) ToJson() io.Reader {
	buffer := new(bytes.Buffer)
	json.NewEncoder(buffer).Encode(u)

	return buffer
}

func UserEntityFromJson(b io.Reader) (*UserEntity, error) {
	user := new(UserEntity)
	json.NewDecoder(b).Decode(user)

	return user, nil
}
