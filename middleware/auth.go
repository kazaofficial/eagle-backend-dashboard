package middleware

import (
	"eagle-backend-dashboard/dto"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

var excludedPaths = []string{"/api/v1/login"}

func AuthenticationMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Check if the request path should be excluded from authentication
		if shouldExcludePath(c) {
			return c.Next()
		}
		token := c.Get("Authorization")
		if token == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(dto.ErrorResponse{
				StatusCode: fiber.StatusUnauthorized,
				Message:    "Unauthorized",
				Error:      "Authorization header not provided",
			})
		}

		splitToken := strings.Split(token, "Bearer ")
		if len(splitToken) != 2 {
			return c.Status(fiber.StatusUnauthorized).JSON(dto.ErrorResponse{
				StatusCode: fiber.StatusUnauthorized,
				Message:    "Unauthorized",
				Error:      "Invalid Authorization header format",
			})
		}

		access_token, err := jwt.ParseWithClaims(splitToken[1], &dto.Claims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("APP_SECRET")), nil
		})

		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(dto.ErrorResponse{
				StatusCode: fiber.StatusUnauthorized,
				Message:    "Unauthorized",
				Error:      err.Error(),
			})
		}

		if Claims, ok := access_token.Claims.(*dto.Claims); ok && access_token.Valid {
			c.Locals("id", Claims.ID)
			c.Locals("username", Claims.Username)
		} else {
			return c.Status(fiber.StatusUnauthorized).JSON(dto.ErrorResponse{
				StatusCode: fiber.StatusUnauthorized,
				Message:    "Unauthorized",
				Error:      "Invalid token",
			})
		}

		return c.Next()
	}
}

// shouldExcludePath checks if the given path should be excluded from authentication
func shouldExcludePath(c *fiber.Ctx) bool {
	path := c.Path()

	for _, excludedPath := range excludedPaths {
		if path == excludedPath {
			return true
		}
	}

	return false
}
