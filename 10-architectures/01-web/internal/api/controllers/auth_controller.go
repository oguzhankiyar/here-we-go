package controllers

import (
	"net/http"
	"time"
	"web-sample/internal/common/models"
	"web-sample/internal/infrastructure/auth"
	configModels "web-sample/internal/infrastructure/config/models"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type AuthController struct {
	appConfig configModels.AppConfig
}

func NewAuthController(appConfig configModels.AppConfig) *AuthController {
	return &AuthController{
		appConfig: appConfig,
	}
}

// HandleToken godoc
// @Summary Get authentication token.
// @Description get authentication token.
// @Tags Auth
// @Accept x-www-form-urlencoded
// @Produce json
// @Param username formData string true "username"
// @Param password formData string true "password"
// @Success 200 {object} models.BaseReponseModel
// @Router /auth/token [post]
func (c *AuthController) HandleToken(ctx echo.Context) error {
	username := ctx.FormValue("username")
	password := ctx.FormValue("password")

	if username != c.appConfig.AdminUsername || password != c.appConfig.AdminPassword {
		return echo.ErrUnauthorized
	}

	expires := time.Now().Add(time.Hour * 72)

	claims := &auth.AuthClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expires.Unix(),
		},
		Name: c.appConfig.AdminUsername,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(c.appConfig.JwtSecret))
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, models.BaseReponseModel{
		Code:    "200",
		Message: "SUCCEEDED",
		Errors:  []string{},
		Data: map[string]interface{}{
			"accessToken": tokenString,
			"tokenType":   "Bearer",
			"expiresIn":   expires.Minute(),
		},
	})
}

// HandleMe godoc
// @Summary Get authenticated user information.
// @Description get authenticated user information.
// @Tags Auth
// @Accept */*
// @Security ApiKeyAuth
// @Produce json
// @Success 200 {object} models.BaseReponseModel
// @Router /auth/me [get]
func (c *AuthController) HandleMe(ctx echo.Context) error {
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(*auth.AuthClaims)
	return ctx.JSON(http.StatusOK, models.BaseReponseModel{
		Code:    "200",
		Message: "SUCCEEDED",
		Errors:  []string{},
		Data: map[string]interface{}{
			"name": claims.Name,
		},
	})
}
