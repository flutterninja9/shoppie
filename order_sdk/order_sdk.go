package ordersdk

import (
	"encoding/json"
	"net/http"
)

type OrderSdk struct {
	baseUrl string
}

func NewOrderSdk(serviceBaseUrl string) OrderSdk {
	return OrderSdk{
		baseUrl: serviceBaseUrl,
	}
}

func (sdk *OrderSdk) GetUserOrders(userId string, authToken string) ([]Order, error) {
	url := sdk.baseUrl + "/user/" + userId

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Add("Authorization", "Bearer "+authToken)

	client := http.DefaultClient
	res, reqErr := client.Do(request)

	if reqErr != nil {
		return nil, reqErr
	}
	defer res.Body.Close()

	var result []Order
	decodeErr := json.NewDecoder(res.Body).Decode(&result)

	if decodeErr != nil {
		return nil, decodeErr
	}

	return result, nil
}
