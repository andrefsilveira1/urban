package routes

import (
	"github.com/andrefsilveira1/urban/internal/transport/rest/endpoints"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/hello", endpoints.Hello)
}
