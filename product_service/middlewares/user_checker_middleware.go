package middlewares

import (
	"os"

	usersdk "github.com/flutterninja9/shoppie/user_sdk"

	"github.com/gofiber/fiber/v2"
)

func UserCheckerMiddleWare(c *fiber.Ctx) error {
	authInfo, ok := c.Locals("authInfo").(*AuthInfo)

	if !ok {
		return c.SendStatus(401)
	}

	userService := usersdk.NewUserService(os.Getenv("USER_SERVICE_BASE_URL"))
	userExists := userService.UserExistsWithId(authInfo.Claims["user_id"].(string), authInfo.Token)
	if !userExists {
		return c.SendStatus(401)
	}

	return c.Next()
}
