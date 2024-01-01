package ordersdk

import "github.com/go-playground/validator/v10"

type OrderSdk struct {
	baseUrl string
}

var Validator = validator.New()

func NewOrderSdk(serviceBaseUrl string) OrderSdk {
	return OrderSdk{
		baseUrl: serviceBaseUrl,
	}
}
