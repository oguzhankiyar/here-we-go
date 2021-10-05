package persistence

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"web-sample/internal/infrastructure/config/models"
)

type PostgresPersistence struct {
	Database *gorm.DB
}

func NewPostgresPersistence(config models.PersistenceConfig) (*PostgresPersistence, error) {
	dsn := fmt.Sprintf("host=%s port=%v user=%s dbname=%s password=%s",
		config.Host,
		config.Port,
		config.User,
		config.DatabaseName,
		config.Password,
	)

	database, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		return nil, err
	}

	return &PostgresPersistence{
		Database: database,
	}, nil
}
