package di

import (
	cartsdk "github.com/flutterninja9/shoppie/cart_sdk"
	ordersdk "github.com/flutterninja9/shoppie/order_sdk"
	productsdk "github.com/flutterninja9/shoppie/product_sdk"
	"github.com/flutterninja9/shoppie/shoppie_bff/config"
	usersdk "github.com/flutterninja9/shoppie/user_sdk"
	"github.com/sirupsen/logrus"
	"go.uber.org/dig"
)

func SetupDi(c *dig.Container) error {
	c.Provide(func() (*config.AppConfig, error) {
		config := config.NewAppConfig()
		return config, nil
	})

	c.Provide(func() (*logrus.Logger, error) {
		logger := logrus.New()
		return logger, nil
	})

	c.Provide(func(c *config.AppConfig) (*usersdk.UserService, error) {
		return usersdk.NewUserService(c.UserServiceBaseUrl), nil
	})

	c.Provide(func(c *config.AppConfig) (productsdk.ProductSdk, error) {
		return productsdk.GetProductSdkInstance(c.ProductServiceBaseUrl), nil
	})

	c.Provide(func(c *config.AppConfig) (ordersdk.OrderSdk, error) {
		return ordersdk.NewOrderSdk(c.OrderServiceBaseUrl), nil
	})

	c.Provide(func(c *config.AppConfig) (cartsdk.CartSdk, error) {
		return cartsdk.NewCartSdk(c.CartServiceBaseUrl), nil
	})

	return nil
}
