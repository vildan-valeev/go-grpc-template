package proto

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"go-bolvanka/internal/config"
	"go-bolvanka/internal/domain/models"
	"go-bolvanka/internal/domain/service"
)

func New(cfg config.Config, cs service.Category, is service.Item) *fiber.App {
	app := fiber.New()

	app.Get("/categories", func(c *fiber.Ctx) error {
		categories, err := cs.GetAll(c.Context())
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				return c.SendStatus(fiber.StatusNotFound)
			}

			return c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.Status(200).JSON(categories)
	})
	app.Get("/categories/:id", func(c *fiber.Ctx) error {
		id, err := uuid.Parse(c.Params("id"))
		if err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		category, err := cs.GetByID(c.Context(), id)
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				return c.SendStatus(fiber.StatusNotFound)
			}

			return c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.Status(200).JSON(category)
	})
	app.Post("/categories", func(c *fiber.Ctx) error {
		var newCategory models.Category
		if err := c.BodyParser(&newCategory); err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}

		err := cs.Create(c.Context(), newCategory)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}

		return c.Status(200).JSON(fiber.Map{"message": "CREATED!"})
	})

	app.Get("/items", func(c *fiber.Ctx) error {
		items, err := is.GetAll(c.Context())
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				return c.SendStatus(fiber.StatusNotFound)
			}

			return c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.Status(200).JSON(items)
	})
	app.Get("/items/:id", func(c *fiber.Ctx) error {
		id, err := uuid.Parse(c.Params("id"))
		if err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		item, err := is.GetByID(c.Context(), id)
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				return c.SendStatus(fiber.StatusNotFound)
			}

			return c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.Status(200).JSON(item)
	})
	app.Post("/items", func(c *fiber.Ctx) error {
		var newItem models.Item
		if err := c.BodyParser(&newItem); err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}

		err := is.Create(c.Context(), newItem)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}

		return c.Status(200).JSON(fiber.Map{"message": "CREATED!"})
	})

	return app
}
