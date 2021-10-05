package interfaces

import (
	"web-sample/internal/infrastructure/config/models"
)

type ConfigParser interface {
	Parse() (*models.Config, error)
}
