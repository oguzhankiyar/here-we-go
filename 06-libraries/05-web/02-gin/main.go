package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var items map[string]string

func main() {
	r := gin.Default()

	r.GET("/", homeHandler)
	r.GET("/items", itemListHandler)
	r.GET("/items/:id", itemDetailHandler)
	r.POST("/items", itemCreateHandler)
	r.PUT("/items/:id", itemUpdateHandler)
	r.DELETE("/items/:id", itemDeleteHandler)

	r.Run()
}

func homeHandler(c *gin.Context) {
	c.String(http.StatusOK, "Hello, World!")
}

func itemListHandler(c *gin.Context) {
	list := make([]map[string]interface{}, 0)

	for id, value := range items {
		list = append(list, map[string]interface{}{
			"id": id,
			"value": value,
		})
	}

	c.JSON(http.StatusOK, list)
}

func itemDetailHandler(c *gin.Context) {
	id := c.Param("id")

	if len(id) == 0 {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "bad request",
		})
		return
	}

	value, exists := items[id]
	if !exists {
		c.JSON(http.StatusNotFound, map[string]interface{}{
			"error": "not found",
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
		"value": value,
	})
}

func itemCreateHandler(c *gin.Context) {
	type Model struct {
		id string
		value string
	}

	var model Model

	if err := c.Bind(&model); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "bad request",
		})
		return
	}

	items[model.id] = model.value

	c.JSON(http.StatusCreated, map[string]interface{}{
		"id": model.id,
		"value": model.value,
	})
}

func itemUpdateHandler(c *gin.Context) {
	id := c.Param("id")

	if len(id) == 0 {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "bad request",
		})
		return
	}

	type Model struct {
		value string
	}

	var model Model

	if err := c.Bind(&model); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "bad request",
		})
		return
	}

	_, exists := items[id]
	if !exists {
		c.JSON(http.StatusNotFound, map[string]interface{}{
			"error": "not found",
		})
		return
	}

	items[id] = model.value

	c.Status(http.StatusNoContent)
}

func itemDeleteHandler(c *gin.Context) {
	id := c.Param("id")

	if len(id) == 0 {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "bad request",
		})
		return
	}

	_, exists := items[id]
	if !exists {
		c.JSON(http.StatusNotFound, map[string]interface{}{
			"error": "not found",
		})
		return
	}

	delete(items, id)

	c.Status(http.StatusNoContent)
}