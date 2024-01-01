package ordersdk

type OrderSdk struct {
	baseUrl string
}

func NewOrderSdk(serviceBaseUrl string) OrderSdk {
	return OrderSdk{
		baseUrl: serviceBaseUrl,
	}
}
