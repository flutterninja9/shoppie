package usersdk

import (
	"net/http"
)

func (r *UserService) GetUserById(id string, authToken string) (*UserEntity, error) {
	requestUrl := r.baseUrl + "/user"
	request, err := http.NewRequest(http.MethodGet, requestUrl, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Add("Authorization", "Bearer "+authToken)

	client := &http.Client{}
	res, resErr := client.Do(request)
	if resErr != nil {
		return nil, resErr
	}

	return UserEntityFromJson(res.Body)
}

// Returns [accessToken] if logged in successfully
func (r *UserService) Login(l *LoginRequest) (*LoginResponse, error) {
	url := r.baseUrl + "/login"
	req, err := http.NewRequest(http.MethodPost, url, l.ToJson())

	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}

	res, resErr := client.Do(req)

	if resErr != nil {
		return nil, resErr
	}

	return LoginResponseFromJson(res.Body)
}

func (r *UserService) Register(l *RegisterRequest) (*UserEntity, error) {
	url := r.baseUrl + "/register"
	req, err := http.NewRequest(http.MethodPost, url, l.ToJson())

	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}

	res, resErr := client.Do(req)

	if resErr != nil {
		return nil, resErr
	}

	return UserEntityFromJson(res.Body)
}

func (rec *UserService) UserExistsWithId(id string, authToken string) bool {
	_, err := rec.GetUserById(id, authToken)
	return err == nil
}
