package middleware

import (
	"github.com/andrefsilveira1/urban/internal/domain"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"
)

func Auth(c *fiber.Ctx) error {
	var secret = []byte("my-scecret-key") // change later
	header := c.Get("Authorization")
	if header == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Missing Token Header",
		})
	}
	tokenString := header[7:]
	claims := domain.Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid token",
		})
	}

	if !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid token",
		})
	}

	return nil
}
