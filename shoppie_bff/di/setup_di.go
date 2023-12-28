package di

import (
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

	return nil
}
