package controllers

import (
	"user_service/database"
	"user_service/middleware"
	"user_service/models"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserRequest struct {
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (u *CreateUserRequest) toDBUser() *models.UserModel {
	user := new(models.UserModel)
	user.Username = u.Username
	user.FirstName = u.FirstName
	user.LastName = u.LastName
	user.Email = u.Email
	user.Password = u.Password

	return user
}

func Register(c *fiber.Ctx, logger *logrus.Logger) error {
	user := new(CreateUserRequest)
	parseErr := c.BodyParser(user)
	if parseErr != nil {
		logger.Warning(parseErr)
		return c.Status(400).JSON(fiber.Map{
			"error": "Please make sure all the fields are correctly filled",
		})
	}

	createdUser, dbErr := database.CreateUser(user.toDBUser())
	if dbErr != nil {
		logger.Warning(dbErr)
		return c.Status(500).JSON(fiber.Map{
			"error": "Something went wrong",
		})
	}

	return c.JSON(createdUser)
}

func Login(c *fiber.Ctx, logger *logrus.Logger) error {
	req := new(LoginRequest)
	if err := c.BodyParser(req); err != nil {
		logger.Error(err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request format",
		})
	}

	user, findErr := database.GetUserByEmail(req.Email)
	if findErr != nil {
		logger.Warning(findErr)
		return c.Status(400).JSON(fiber.Map{
			"error": "Please cross check the entered details",
		})
	}

	passwordsMatch := models.PasswordsMatch(req.Password, user.Password)
	token, tokenGenErr := user.GenerateToken()

	if tokenGenErr != nil {
		logger.Warning(tokenGenErr)
		return c.Status(500).JSON(fiber.Map{
			"error": "Something went wrong while generating token",
		})
	}

	if !passwordsMatch {
		logger.Info("Passwords did not match")
		return c.Status(400).JSON(fiber.Map{
			"error": "Please cross check the entered details",
		})

	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Logged in successfully",
		"token":   token,
	})
}

func GetUserById(c *fiber.Ctx, logger *logrus.Logger) error {
	data, ok := c.Locals("authInfo").(*middleware.AuthInfo)
	if !ok {
		return c.Status(401).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}
	user, err := database.GetUserById(data.Claims["user_id"].(string))

	if err != nil {
		logger.Warning(err)
		return c.Status(400).JSON(fiber.Map{
			"error": "No such user found",
		})
	}

	return c.Status(200).JSON(user)
}
