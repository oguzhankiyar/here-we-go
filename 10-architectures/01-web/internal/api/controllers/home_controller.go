package controllers

import (
	"net/http"
	"strconv"
	"web-sample/internal/infrastructure/config/models"

	"github.com/labstack/echo/v4"
)

type HomeController struct {
	appConfig models.AppConfig
}

func NewHomeController(appConfig models.AppConfig) *HomeController {
	return &HomeController{
		appConfig: appConfig,
	}
}

// Handle godoc
// @Summary Show the information of server.
// @Description get the information of server.
// @Tags Home
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router / [get]
func (c *HomeController) Handle(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"name":    c.appConfig.Name,
		"env":     c.appConfig.Environment,
		"version": c.appConfig.Version,
		"host":    c.appConfig.Host,
		"port":    c.appConfig.Port,
		"docs":    c.appConfig.Host + ":" + strconv.Itoa(c.appConfig.Port) + "/swagger",
	})
}
