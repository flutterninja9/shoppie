package config

import "os"

type AppConfig struct {
	UserServiceBaseUrl    string
	ProductServiceBaseUrl string
	OrderServiceBaseUrl   string
	CartServiceBaseUrl    string
	ServerPort            string
}

func NewAppConfig() *AppConfig {
	return &AppConfig{
		UserServiceBaseUrl:    os.Getenv("USER_SERVICE_BASE_URL"),
		ProductServiceBaseUrl: os.Getenv("PRODUCT_SERVICE_BASE_URL"),
		OrderServiceBaseUrl:   os.Getenv("ORDER_SERVICE_BASE_URL"),
		CartServiceBaseUrl:    os.Getenv("CART_SERVICE_BASE_URL"),
		ServerPort:            os.Getenv("SERVER_PORT"),
	}
}
