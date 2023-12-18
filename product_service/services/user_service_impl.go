package services

import (
	"encoding/json"
	"net/http"
	"os"
	"sync"
)

var (
	once    sync.Once
	service *UserServiceImpl
)

type UserServiceImpl struct{}

func GetUserService() *UserServiceImpl {
	once.Do(func() {
		service = &UserServiceImpl{}
	})

	return service
}

func (rec *UserServiceImpl) GetUserById(id string, authToken string) (*UserEntity, error) {
	var usersApiBaseUrl = os.Getenv("USER_SERVICE_BASE_URL")
	requestUrl := usersApiBaseUrl + "/user"
	request, err := http.NewRequest("GET", requestUrl, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Add("Authorization", "Bearer "+authToken)

	client := &http.Client{}
	res, resErr := client.Do(request)
	if resErr != nil {
		return nil, resErr
	}

	var user UserEntity
	err = json.NewDecoder(res.Body).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (rec *UserServiceImpl) UserExistsWithId(id string, authToken string) bool {
	_, err := rec.GetUserById(id, authToken)
	return err == nil
}
