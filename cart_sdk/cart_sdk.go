package cartsdk

type CartSdk struct {
	baseUrl string
}

func NewCartSdk(baseUrl string) *CartSdk {
	return &CartSdk{baseUrl: baseUrl}
}
