package ordersdk

import (
	"net/http"
)

func (sdk *OrderSdk) GetUserOrders(userId string, authToken string) (*Orders, error) {
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

	return OrdersFromJson(res.Body)
}
