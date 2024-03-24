package rest

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func Start(port int) error {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})

	return app.Listen(":" + strconv.Itoa(port))
}
