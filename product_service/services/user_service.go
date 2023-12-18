package services

type IUserService interface {
	GetUserById(id string, authToken string) (*UserEntity, error)
	UserExistsWithId(id string, authToken string) bool
}
