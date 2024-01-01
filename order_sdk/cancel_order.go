package ordersdk

import "net/http"

func (o *OrderSdk) CancelOrder(id string, token string) (*CancelOrderResponse, error) {
	url := o.baseUrl + "/" + id + "/cancel"
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

	return CancelOrderResponseToJson(res.Body)
}
