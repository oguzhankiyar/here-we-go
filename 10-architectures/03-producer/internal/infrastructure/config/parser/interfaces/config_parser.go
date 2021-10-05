package interfaces

import (
	"producer-sample/internal/infrastructure/config/models"
)

type ConfigParser interface {
	Parse() (*models.Config, error)
}
