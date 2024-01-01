package ordersdk

import "net/http"

func (s *OrderSdk) GetAllOrders(token string) (*Orders, error) {
	url := s.baseUrl + "/"
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

	return OrdersFromJson(res.Body)
}
