package handler

import (
	"github.com/flutterninja9/shoppie/log_service/db"
	"github.com/flutterninja9/shoppie/log_service/model"
	"github.com/gofiber/fiber/v2"
)

func HandleSaveLogs(c *fiber.Ctx) error {
	log := new(model.LogModel)
	err := c.BodyParser(log)
	if err != nil {
		return c.SendStatus(fiber.ErrBadRequest.Code)
	}

	saveErr := log.Save(db.DBClient)
	if saveErr != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": saveErr})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Log saved"})
}
