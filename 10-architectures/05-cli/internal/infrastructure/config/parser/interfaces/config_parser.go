package interfaces

import (
	"cli-sample/internal/infrastructure/config/models"
)

type ConfigParser interface {
	Parse() (*models.Config, error)
}
