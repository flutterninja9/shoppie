package ordersdk

import (
	"fmt"
	"net/http"
)

func (o *OrderSdk) CancelOrder(id string, token string) (bool, error) {
	url := o.baseUrl + "/" + id + "/cancel"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	fmt.Println(url)
	if err != nil {
		return false, err
	}

	req.Header.Add("Authorization", "Bearer "+token)
	client := http.Client{}
	res, err := client.Do(req)

	if err != nil {
		return false, err
	}

	if res.StatusCode != http.StatusOK {
		return false, nil
	}

	return true, nil
}
