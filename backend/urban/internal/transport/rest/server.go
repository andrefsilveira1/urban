package rest

import (
	"strconv"

	"github.com/andrefsilveira1/urban/internal/domain"
	"github.com/andrefsilveira1/urban/internal/transport/rest/routes"
	"github.com/gofiber/fiber/v2"
)

func Start(port int, userService *domain.UserService) error {
	app := fiber.New()
	routes.Setup(app, userService)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})

	return app.Listen(":" + strconv.Itoa(port))
}
