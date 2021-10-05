package controllers

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type HealthController struct {
}

func NewHealthController() *HealthController {
	return &HealthController{}
}

// Handle godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags Health
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /health [get]
func (c *HealthController) Handle(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"time": time.Now().Unix(),
	})
}
