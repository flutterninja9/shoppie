package usersdk

import (
	"encoding/json"
	"net/http"
)

func (rec *UserService) GetUserById(id string, authToken string) (*UserEntity, error) {
	requestUrl := rec.baseUrl + "/user"
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

func (rec *UserService) UserExistsWithId(id string, authToken string) bool {
	_, err := rec.GetUserById(id, authToken)
	return err == nil
}
