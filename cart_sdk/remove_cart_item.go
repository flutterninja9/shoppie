package cartsdk

import (
	"net/http"
)

func (s *CartSdk) RemoveCartItem(userId string, itemId string, token string) (*CommonCartRespose, error) {
	url := s.baseUrl + "/" + userId + "/items" + "/" + itemId

	req, err := http.NewRequest(http.MethodDelete, url, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+token)
	client := http.Client{}
	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	return CommonCartResposeFromJson(res.Body)
}
