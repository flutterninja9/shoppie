package cartsdk

import (
	"net/http"
)

func (s *CartSdk) UpdateCartItem(u *UpdateCartItemRequest, userId string, itemId string, token string) (*CommonCartRespose, error) {
	url := s.baseUrl + "/" + userId + "/items" + "/" + itemId
	json, err := u.ToJson()

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPut, url, json)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("Content-Type", "application/json")
	client := http.Client{}
	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	return CommonCartResposeFromJson(res.Body)
}
