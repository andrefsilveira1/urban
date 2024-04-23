package rest

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/andrefsilveira1/urban/internal/domain"
	"github.com/andrefsilveira1/urban/internal/transport/rest/routes"
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	app *fiber.App
}

func Start(port int, userService *domain.UserService) error {
	app := fiber.New()
	routes.Setup(app, userService)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})

	return app.Listen(":" + strconv.Itoa(port))
}

func (s *Server) Stop(timeout time.Duration) error {
	_, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	if err := s.app.Shutdown(); err != nil {
		return fmt.Errorf("failed to gracefully shutdown server: %v", err)
	}

	return nil
}
