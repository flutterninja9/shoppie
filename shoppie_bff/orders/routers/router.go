package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func SetupRouters(l *logrus.Logger, a *fiber.App) error {
	l.Info("Setting up order router")
	return nil
}
