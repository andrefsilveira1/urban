package endpoints

import (
	"github.com/andrefsilveira1/urban/internal/domain"
	"github.com/andrefsilveira1/urban/internal/domain/models"
	"github.com/gofiber/fiber"
)

func HelloUser(c *fiber.Ctx, userService *domain.UserService) error {
	return c.JSON("Hello user")
}

func Register(c *fiber.Ctx, userService *domain.UserService) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	err := userService.Register(user.Name string, user.Email string, user.Password string)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	return c.JSON("User created")
}
