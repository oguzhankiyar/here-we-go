package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gavv/httpexpect/v2"
	"github.com/labstack/echo"
)

func Test(t *testing.T) {
	handler := echo.New()
	handler.GET("/items", func(c echo.Context) error {
		c.JSON(200, []struct{}{})
		return nil
	})

	server := httptest.NewServer(handler)
	defer server.Close()

	e := httpexpect.New(t, server.URL)

	e.GET("/items").
		Expect().
		Status(http.StatusOK).JSON().Array().Empty()
}