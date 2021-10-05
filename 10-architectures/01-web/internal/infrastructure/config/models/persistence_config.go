package models

import "web-sample/internal/infrastructure/config/errors"

type PersistenceConfig struct {
	Host         string `mapstructure:"PERSISTENCE_HOST" env:"PERSISTENCE_HOST"`
	Port         int    `mapstructure:"PERSISTENCE_PORT" env:"PERSISTENCE_PORT"`
	SSLMode      bool   `mapstructure:"PERSISTENCE_SSL_MODE" env:"PERSISTENCE_SSL_MODE"`
	User         string `mapstructure:"PERSISTENCE_USER" env:"PERSISTENCE_USER"`
	Password     string `mapstructure:"PERSISTENCE_PASSWORD" env:"PERSISTENCE_PASSWORD"`
	DatabaseName string `mapstructure:"PERSISTENCE_DATABASE_NAME" env:"PERSISTENCE_DATABASE_NAME"`
}

func NewPersistenceConfig() *PersistenceConfig {
	return &PersistenceConfig{}
}

func (c *PersistenceConfig) Validate() error {
	if len(c.Host) == 0 {
		return errors.NewConfigValidateError("invalid persistence host")
	}

	if c.Port == 0 {
		return errors.NewConfigValidateError("invalid persistence port")
	}

	if len(c.User) == 0 {
		return errors.NewConfigValidateError("invalid persistence user")
	}

	if len(c.Password) == 0 {
		return errors.NewConfigValidateError("invalid persistence password")
	}

	if len(c.DatabaseName) == 0 {
		return errors.NewConfigValidateError("invalid persistence database name")
	}

	return nil
}
