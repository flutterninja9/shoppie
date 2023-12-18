package middlewares

import (
	"product_service/services"

	"github.com/gofiber/fiber/v2"
)

func UserCheckerMiddleWare(c *fiber.Ctx) error {
	authInfo, ok := c.Locals("authInfo").(*AuthInfo)

	if !ok {
		return c.SendStatus(401)
	}

	userExists := services.GetUserService().UserExistsWithId(authInfo.Claims["user_id"].(string), authInfo.Token)
	if !userExists {
		return c.SendStatus(401)
	}

	return c.Next()
}
