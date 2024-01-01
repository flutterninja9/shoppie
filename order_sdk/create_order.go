package ordersdk

import "net/http"

func (s *OrderSdk) CreateOrder(r *CreateOrderRequest, token string) (*Order, error) {
	url := s.baseUrl + "/"
	body, err := r.ToJson()

	if err != nil {
		return nil, err
	}

	validationErr := Validator.Struct(body)

	if validationErr != nil {
		return nil, validationErr
	}
	req, err := http.NewRequest(http.MethodPost, url, body)

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
