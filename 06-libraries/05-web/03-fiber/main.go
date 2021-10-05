package main

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

var items map[string]string

func main() {
	app := fiber.New()

	app.Get("/", homeHandler)
	app.Get("/items", itemListHandler)
	app.Get("/items/:id", itemDetailHandler)
	app.Post("/items", itemCreateHandler)
	app.Put("/items/:id", itemUpdateHandler)
	app.Delete("/items/:id", itemDeleteHandler)

	app.Listen(":3000")
}

func homeHandler(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).SendString("Hello, World!")
}

func itemListHandler(c *fiber.Ctx) error {
	list := make([]map[string]interface{}, 0)

	for id, value := range items {
		list = append(list, map[string]interface{}{
			"id": id,
			"value": value,
		})
	}

	return c.Status(http.StatusOK).JSON(list)
}

func itemDetailHandler(c *fiber.Ctx) error {
	id := c.Params("id")

	if len(id) == 0 {
		return c.Status(http.StatusBadRequest).JSON(map[string]interface{}{
			"error": "bad request",
		})
	}

	value, exists := items[id]
	if !exists {
		return c.Status(http.StatusNotFound).JSON(map[string]interface{}{
			"error": "not found",
		})
	}

	return c.Status(http.StatusOK).JSON(map[string]interface{}{
		"id": id,
		"value": value,
	})
}

func itemCreateHandler(c *fiber.Ctx) error {
	type Model struct {
		id string
		value string
	}

	var model Model

	if err := c.BodyParser(&model); err != nil {
		return c.Status(http.StatusBadRequest).JSON(map[string]interface{}{
			"error": "bad request",
		})
	}

	items[model.id] = model.value

	return c.Status(http.StatusCreated).JSON(map[string]interface{}{
		"id": model.id,
		"value": model.value,
	})
}

func itemUpdateHandler(c *fiber.Ctx) error {
	id := c.Params("id")

	if len(id) == 0 {
		return c.Status(http.StatusBadRequest).JSON(map[string]interface{}{
			"error": "bad request",
		})
	}

	type Model struct {
		value string
	}

	var model Model

	if err := c.BodyParser(&model); err != nil {
		return c.Status(http.StatusBadRequest).JSON(map[string]interface{}{
			"error": "bad request",
		})
	}

	_, exists := items[id]
	if !exists {
		return c.Status(http.StatusNotFound).JSON(map[string]interface{}{
			"error": "not found",
		})
	}

	items[id] = model.value

	return c.Status(http.StatusNoContent).Send(nil)
}

func itemDeleteHandler(c *fiber.Ctx) error {
	id := c.Params("id")

	if len(id) == 0 {
		return c.Status(http.StatusBadRequest).JSON(map[string]interface{}{
			"error": "bad request",
		})
	}

	_, exists := items[id]
	if !exists {
		return c.Status(http.StatusNotFound).JSON(map[string]interface{}{
			"error": "not found",
		})
	}

	delete(items, id)

	return c.Status(http.StatusNoContent).Send(nil)
}