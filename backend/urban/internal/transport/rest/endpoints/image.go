package endpoints

import (
	"github.com/andrefsilveira1/urban/internal/domain"
	"github.com/andrefsilveira1/urban/internal/domain/entity"
	"github.com/gofiber/fiber/v2"
)

func Hello(c *fiber.Ctx) error {
	return c.JSON("Hello")
}

func Save(c *fiber.Ctx, imageService *domain.ImageService) error {
	var image entity.Image
	if err := c.BodyParser(&image); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request body")
	}
	err := imageService.Register(image.Name, image.Date, image.Content)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON("Image saved")
}

func Get(c *fiber.Ctx, imageService *domain.ImageService) error {
	id := c.Params("id")
	image, err := imageService.Get(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(image)
}

func List(c *fiber.Ctx, imageService *domain.ImageService) error {
	images, err := imageService.List()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(images)
}

func Delete(c *fiber.Ctx, imageService *domain.ImageService) error {
	id := c.Params("id")
	err := imageService.Delete(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON("Image deleted...")
}
