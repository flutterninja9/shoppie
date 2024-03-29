package usersdk

import (
	"log"
	"net/http"
)

func (r *UserService) GetUserById(id string, authToken string) (*UserEntity, error) {
	requestUrl := r.b + "/user"
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
	url := r.b + "/login"
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

	if res.StatusCode != http.StatusOK {
		return nil, resErr
	}

	return LoginResponseFromJson(res.Body)
}

func (r *UserService) Register(l *RegisterRequest) (*UserEntity, error) {
	log.Println("Contructing register url")
	url := r.b + "/register"
	log.Println("Contructed url", url)
	log.Println("Trying to call user-service register api")
	req, err := http.NewRequest(http.MethodPost, url, l.ToJson())

	if err != nil {
		log.Fatal("Error in register api req construction", err)
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}

	res, resErr := client.Do(req)
	log.Println("Got response from user-service register api with statusCode", res.StatusCode, resErr)

	if resErr != nil {
		return nil, resErr
	}

	return UserEntityFromJson(res.Body)
}

func (rec *UserService) UserExistsWithId(id string, authToken string) bool {
	_, err := rec.GetUserById(id, authToken)
	return err == nil
}

func (u *UserService) GetOrders(accessToken string) (*Orders, error) {
	url := u.b + "/orders"

	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+accessToken)

	client := http.Client{}
	res, resErr := client.Do(req)

	if resErr != nil {
		return nil, resErr
	}

	return OrdersFromJson(res.Body), nil
}
