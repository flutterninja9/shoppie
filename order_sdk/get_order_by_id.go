package ordersdk

import "net/http"

func (s *OrderSdk) GetOrderById(id string, token string) (*Order, error) {
	url := s.baseUrl + "/" + id
	req, err := http.NewRequest(http.MethodPost, url, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+token)
	
	client := http.Client{}
	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}
	return OrderFromJson(res.Body)
}
