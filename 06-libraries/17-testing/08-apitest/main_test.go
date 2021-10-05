package main

import (
	"net/http"
	"testing"

	"github.com/labstack/echo"
	"github.com/steinfletcher/apitest"
)

func Test(t *testing.T) {
	handler := echo.New()
	handler.GET("/items", func(c echo.Context) error {
		c.JSON(200, []interface {} {
			map[string]interface {} {
				"id": 10,
				"name": "Gopher",
			},
		})
		return nil
	})

	apitest.New().
		Handler(handler).
		Get("/items").
		Expect(t).
		Body(`[{"id": 10, "name": "Gopher"}]`).
		Status(http.StatusOK).
		End()
}