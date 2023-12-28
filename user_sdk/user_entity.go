package usersdk

import "github.com/google/uuid"

type UserEntity struct {
	ID    uuid.UUID `json:"id"`
	Email string    `json:"email"`
}
