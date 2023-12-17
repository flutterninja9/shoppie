package middleware

import (
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

// AuthInfo struct to store JWT token and its claims
type AuthInfo struct {
	Token  string
	Claims jwt.MapClaims
}

// AuthMiddleware is a Fiber middleware for extracting and decoding the JWT token from the Authorization header
func AuthMiddleware(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	bearerToken := ""

	if strings.HasPrefix(authHeader, "Bearer ") {
		bearerToken = strings.TrimPrefix(authHeader, "Bearer ")
	}

	// Decode the token
	token, err := jwt.Parse(bearerToken, func(token *jwt.Token) (interface{}, error) {
		// Make sure token's method matches the signing method you expect
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fiber.NewError(fiber.StatusUnauthorized, "Unexpected signing method")
		}
		// Use the secret key used to sign the token
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid token")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		authInfo := &AuthInfo{
			Token:  bearerToken,
			Claims: claims,
		}
		c.Locals("authInfo", authInfo)
	} else {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid token")
	}

	return c.Next()
}
