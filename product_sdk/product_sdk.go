package productsdk

type ProductSdk struct {
	BaseUrl string
}

func GetProductSdkInstance(baseUrl string) ProductSdk {
	return ProductSdk{
		BaseUrl: baseUrl,
	}
}
