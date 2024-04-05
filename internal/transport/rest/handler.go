package rest

import (
	"net/http"

	"github.com/Alieksieiev0/user-service/internal/models"
	"github.com/Alieksieiev0/user-service/internal/services"
	"github.com/gofiber/fiber/v2"
)

func getById(service services.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		entity, err := service.GetById(c.Context(), c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"error": err.Error()})
		}

		return c.Status(http.StatusOK).JSON(entity)
	}
}

func create(service services.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := &models.User{}
		if err := c.BodyParser(user); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		if err := service.Save(c.Context(), user); err != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"error": err.Error()})
		}

		return c.Status(fiber.StatusCreated).JSON(user)
	}
}
