package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

var items map[string]string

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", homeHandler)
	e.GET("/items", itemListHandler)
	e.GET("/items/:id", itemDetailHandler)
	e.POST("/items", itemCreateHandler)
	e.PUT("/items/:id", itemUpdateHandler)
	e.DELETE("/items/:id", itemDeleteHandler)

	e.Logger.Fatal(e.Start(":1323"))
}

func homeHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func itemListHandler(c echo.Context) error {
	list := make([]map[string]interface{}, 0)

	for id, value := range items {
		list = append(list, map[string]interface{}{
			"id": id,
			"value": value,
		})
	}

	return c.JSON(http.StatusOK, list)
}

func itemDetailHandler(c echo.Context) error {
	id := c.Param("id")

	if len(id) == 0 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "bad request",
		})
	}

	value, exists := items[id]
	if !exists {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"error": "not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
		"value": value,
	})
}

func itemCreateHandler(c echo.Context) error {
	type Model struct {
		id string
		value string
	}

	var model Model

	if err := c.Bind(&model); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "bad request",
		})
	}

	items[model.id] = model.value

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"id": model.id,
		"value": model.value,
	})
}

func itemUpdateHandler(c echo.Context) error {
	id := c.Param("id")

	if len(id) == 0 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "bad request",
		})
	}

	type Model struct {
		value string
	}

	var model Model

	if err := c.Bind(&model); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "bad request",
		})
	}

	_, exists := items[id]
	if !exists {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"error": "not found",
		})
	}

	items[id] = model.value

	return c.NoContent(http.StatusNoContent)
}

func itemDeleteHandler(c echo.Context) error {
	id := c.Param("id")

	if len(id) == 0 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "bad request",
		})
	}

	_, exists := items[id]
	if !exists {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"error": "not found",
		})
	}

	delete(items, id)

	return c.NoContent(http.StatusNoContent)
}