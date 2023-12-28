package di

import (
	"github.com/flutterninja9/shoppie/shoppie_bff/config"
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

	return nil
}
