package middleware

import (
	"fmt"

	"github.com/gofiber/fiber"
)

func Auth(c *fiber.Ctx) error {
	header := c.Get("Authorization")
	fmt.Println("HEADER:", header)
	// TODO

	return c.Next()
}
