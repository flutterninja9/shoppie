package cartsdk

import "net/http"

func (s *CartSdk) GetUserCart(userId string, token string) (*CartModel, error) {
	url := s.baseUrl + "/" + userId
	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+token)
	client := http.Client{}
	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	return CartModelToJson(res.Body)
}
