package handler

import (
	"github.com/flutterninja9/shoppie/log_service/db"
	"github.com/flutterninja9/shoppie/log_service/model"
	"github.com/gofiber/fiber/v2"
)

func HandleFetchLogs(c *fiber.Ctx) error {
	logs := new(model.Logs)
	logs.GetLogs(db.DBClient)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": logs})
}
