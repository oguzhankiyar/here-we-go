package interfaces

import (
	"cron-sample/internal/infrastructure/config/models"
)

type ConfigParser interface {
	Parse() (*models.Config, error)
}
