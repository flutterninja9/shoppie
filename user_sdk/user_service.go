package usersdk

type UserService struct {
	baseUrl string
}

func NewUserService(baseUrl string) *UserService {
	return &UserService{
		baseUrl: baseUrl,
	}
}