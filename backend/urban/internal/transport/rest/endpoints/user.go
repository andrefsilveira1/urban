package endpoints

import (
	"github.com/andrefsilveira1/urban/internal/domain"
	"github.com/andrefsilveira1/urban/internal/domain/entity"
	"github.com/gofiber/fiber"
)

func HelloUser(c *fiber.Ctx, userService *domain.UserService) error {
	return c.JSON("Hello user")
}

func Register(c *fiber.Ctx, userService *domain.UserService) error {
	var user entity.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	err := userService.Register(user.Name, user.Email, user.Password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	return c.JSON("User created")
}

func GetUser(c *fiber.Ctx, userService *domain.UserService) error {
	id := c.Params("id")
	user, err := userService.Get(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(user)

}

func ListUsers(c *fiber.Ctx, userService *domain.UserService) error {
	users, err := userService.List()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(users)
}
