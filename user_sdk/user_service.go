package usersdk

type UserService struct {
	b string
}

func NewUserService(b string) *UserService {
	return &UserService{
		b: b,
	}
}
