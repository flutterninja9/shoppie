package cartsdk

import (
	"net/http"
)

func (s *CartSdk) ClearCart(userId string, token string) (*CommonCartRespose, error) {
	url := s.baseUrl + "/" + userId

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
