package middlewares

import (
	"os"

	usersdk "github.com/flutterninja9/shoppie/user_sdk"
	"github.com/gofiber/fiber/v2"
)

func UserExistenceCheckMiddleware(c *fiber.Ctx) error {
	authInfo, ok := c.Locals("authInfo").(*AuthInfo)

	if !ok {
		return c.SendStatus(fiber.ErrBadRequest.Code)
	}

	userId := authInfo.Claims["user_id"].(string)
	userService := usersdk.NewUserService(os.Getenv("USER_SERVICE_BASE_URL"))
	userExists := userService.UserExistsWithId(userId, authInfo.Token)

	if !userExists {
		return c.SendStatus(401)
	}

	return c.Next()
}
