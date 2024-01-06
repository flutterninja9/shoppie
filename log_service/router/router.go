package router

import (
	"github.com/flutterninja9/shoppie/log_service/handler"
	"github.com/gofiber/fiber/v2"
)

func SetupRouter(app *fiber.App) {
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	app.Get("/", handler.HandleFetchLogs)
	app.Post("/", handler.HandleSaveLogs)
}
