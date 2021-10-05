package interfaces

import (
	"consumer-sample/internal/infrastructure/config/models"
)

type ConfigParser interface {
	Parse() (*models.Config, error)
}
