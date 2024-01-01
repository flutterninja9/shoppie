package ordersdk

import "net/http"

// Returns updated order if order is successfully updated
func (s *OrderSdk) UpdateOrder(id string, u *UpdateOrderStatusRequest, token string) (*Order, error) {
	url := s.baseUrl + "/" + id
	body, err := u.ToJson()

	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPut, url, body)

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

	return OrderFromJson(res.Body)
}
