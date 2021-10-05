package middlewares

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"web-sample/internal/common/models"
	appErrors "web-sample/internal/infrastructure/errors"
)

func Error() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			err := next(ctx)

			if err == nil {
				return err
			}

			if errors.Is(err, middleware.ErrJWTMissing) || errors.Is(err, middleware.ErrJWTInvalid) || errors.Is(err, echo.ErrUnauthorized) {
				ctx.JSON(http.StatusUnauthorized, models.BaseReponseModel{
					Code:    "401",
					Message: "FAILED",
					Errors: []string{
						"Unauthorized",
					},
				})
				return nil
			}

			switch e := err.(type) {
			case *appErrors.NotFoundError:
				ctx.JSON(http.StatusNotFound, models.BaseReponseModel{
					Code:    e.Code,
					Message: "FAILED",
					Errors: []string{
						e.Error(),
					},
				})
				return nil
			case *appErrors.BadRequestError:
				ctx.JSON(http.StatusBadRequest, models.BaseReponseModel{
					Code:    e.Code,
					Message: "FAILED",
					Errors: []string{
						err.Error(),
					},
				})
				return nil
			case *appErrors.AlreadyExistError:
				ctx.JSON(http.StatusConflict, models.BaseReponseModel{
					Code:    e.Code,
					Message: "FAILED",
					Errors: []string{
						err.Error(),
					},
				})
				return nil
			default:
				ctx.JSON(http.StatusInternalServerError, models.BaseReponseModel{
					Code:    "500",
					Message: "FAILED",
					Errors: []string{
						err.Error(),
					},
				})
				return nil
			}
		}
	}
}
