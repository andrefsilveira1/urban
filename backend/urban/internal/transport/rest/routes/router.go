package routes

import (
	"github.com/andrefsilveira1/urban/internal/domain"
	"github.com/andrefsilveira1/urban/internal/transport/rest/endpoints"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App, userService *domain.UserService, imageService *domain.ImageService) {
	app.Get("/hello", endpoints.Hello)
	app.Post("/user", func(c *fiber.Ctx) error {
		return endpoints.Register(c, userService)
	})
}
